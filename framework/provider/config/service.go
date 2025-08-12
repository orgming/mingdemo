package config

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-yaml/yaml"
	"github.com/mitchellh/mapstructure"
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// MingConfig  配置文件服务
type MingConfig struct {
	c        framework.Container // 容器
	folder   string              // 文件夹
	keyDelim string              // 路径的分隔符，默认为点
	lock     sync.RWMutex        // 配置文件读写锁
	envMaps  map[string]string   // 所有的环境变量
	confMaps map[string]any      // 配置文件结构，key为文件名
	confRaws map[string][]byte   // 配置文件的原始信息
}

func NewMingConfig(params ...any) (any, error) {
	container := params[0].(framework.Container)
	envFolder := params[1].(string)
	envMaps := params[2].(map[string]string)

	if _, err := os.Stat(envFolder); os.IsNotExist(err) {
		return nil, errors.New("folder " + envFolder + " not exist: " + err.Error())
	}

	config := &MingConfig{
		c:        container,
		folder:   envFolder,
		keyDelim: ".",
		envMaps:  envMaps,
		confMaps: make(map[string]any),
		confRaws: make(map[string][]byte),
		lock:     sync.RWMutex{},
	}

	// 读取每个文件
	files, err := os.ReadDir(envFolder)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, file := range files {
		fileName := file.Name()
		err := config.loadConfigFile(envFolder, fileName)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	// 监控文件夹文件
	watch, err := fsnotify.NewWatcher() // TODO
	if err != nil {
		return nil, err
	}
	err = watch.Add(envFolder)
	if err != nil {
		return nil, err
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		for {
			select {
			case ev := <-watch.Events:
				{
					// 判断事件发生的类型
					// Create 创建
					// Write 写入
					// Remove 删除
					path, _ := filepath.Abs(ev.Name)
					index := strings.LastIndex(path, string(os.PathSeparator))
					folder := path[:index]
					fileName := path[index+1:]

					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name)
						config.loadConfigFile(folder, fileName)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name)
						config.loadConfigFile(folder, fileName)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name)
						config.removeConfigFile(folder, fileName)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()

	return config, nil
}

func (conf *MingConfig) IsExist(key string) bool {
	return conf.find(key) != nil
}

func (conf *MingConfig) Get(key string) any {
	return conf.find(key)
}

func (conf *MingConfig) GetBool(key string) bool {
	return cast.ToBool(conf.find(key))
}

func (conf *MingConfig) GetInt(key string) int {
	return cast.ToInt(conf.find(key))
}

func (conf *MingConfig) GetFloat64(key string) float64 {
	return cast.ToFloat64(conf.find(key))
}

func (conf *MingConfig) GetTime(key string) time.Time {
	return cast.ToTime(conf.find(key))
}

func (conf *MingConfig) GetString(key string) string {
	return cast.ToString(conf.find(key))
}

func (conf *MingConfig) GetIntSlice(key string) []int {
	return cast.ToIntSlice(conf.find(key))

}

func (conf *MingConfig) GetStringSlice(key string) []string {
	return cast.ToStringSlice(conf.find(key))

}

func (conf *MingConfig) GetStringMap(key string) map[string]any {
	return cast.ToStringMap(conf.find(key))

}

func (conf *MingConfig) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(conf.find(key))

}

func (conf *MingConfig) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(conf.find(key))

}

func (conf *MingConfig) Load(key string, val any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "yaml",
		Result:  val,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(conf.find(key))
}

// 表示使用环境变量maps替换context中的env(xxx)的环境变量
func replace(content []byte, maps map[string]string) []byte {
	if maps == nil {
		return content
	}
	// 直接使用ReplaceAll替换。这个性能可能不是最优，但是配置文件加载，频率是比较低的，可以接受
	for key, val := range maps {
		reKey := "env(" + key + ")"
		content = bytes.ReplaceAll(content, []byte(reKey), []byte(val))
	}
	return content
}

// 查找某个路径的配置项
func searchMap(source map[string]any, path []string) any {
	if len(path) == 0 {
		return source
	}

	// 判断是否有下个路径
	next, ok := source[path[0]]
	if ok {
		// 判断这个路径是否为1
		if len(path) == 1 {
			return next
		}

		// 判断下一个路径的类型
		switch next.(type) {
		case map[any]any:
			// 如果是interface的map，使用cast进行下value转换
			return searchMap(cast.ToStringMap(next), path[1:])
		case map[string]any:
			// 如果是map[string]，直接循环调用
			return searchMap(next.(map[string]any), path[1:])
		default:
			// 否则的话，返回nil
			return nil
		}
	}
	return nil
}

// 读取某个配置文件
func (conf *MingConfig) loadConfigFile(folder string, file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()

	//  判断文件是否以yaml或者yml作为后缀
	s := strings.Split(file, ".")
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		name := s[0]

		// 读取文件内容
		bf, err := os.ReadFile(filepath.Join(folder, file))
		if err != nil {
			return err
		}
		// 直接针对文本做环境变量的替换
		bf = replace(bf, conf.envMaps)
		// 解析对应的文件
		c := map[string]any{}
		if err := yaml.Unmarshal(bf, &c); err != nil {
			return err
		}
		conf.confMaps[name] = c
		conf.confRaws[name] = bf

		// 读取app.path中的信息，更新app对应的folder
		if name == "app" && conf.c.IsBind(contract.AppKey) {
			if p, ok := c["path"]; ok {
				appService := conf.c.MustMake(contract.AppKey).(contract.App)
				appService.LoadAppConfig(cast.ToStringMapString(p))
			}
		}
	}
	return nil
}

// 删除文件的操作
func (conf *MingConfig) removeConfigFile(folder string, file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	s := strings.Split(file, ".")
	// 只有yaml或者yml后缀才执行
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		name := s[0]
		// 删除内存中对应的key
		delete(conf.confRaws, name)
		delete(conf.confMaps, name)
	}
	return nil
}

func (conf *MingConfig) find(key string) any {
	conf.lock.RLock()
	defer conf.lock.RUnlock()
	return searchMap(conf.confMaps, strings.Split(key, conf.keyDelim))
}
