package env

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/orgming/ming/framework/contract"
	"io"
	"os"
	"path"
	"strings"
)

// MingEnv 是Env的具体实现
type MingEnv struct {
	folder string            // 代表.env所在的目录
	maps   map[string]string // 保存所有的环境变量
}

func NewMingEnv(params ...any) (any, error) {
	if len(params) != 0 {
		return nil, errors.New("NewMingEnv param error")
	}
	folder := params[0].(string)

	env := &MingEnv{
		folder: folder,
		maps:   map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	// 解析folder/.env文件
	file := path.Join(folder, ".env") // TODO 对比 filepath

	f, err := os.Open(file)
	if err == nil {
		defer f.Close()

		br := bufio.NewReader(f)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF { // TODO
				break
			}
			s := bytes.SplitN(line, []byte("="), 2)
			if len(s) < 2 {
				continue
			}
			env.maps[string(s[0])] = string(s[1])
		}
	}

	// 获取当前程序的环境变量，并且覆盖.env文件下的变量
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		env.maps[pair[0]] = pair[1]
	}
	return env, nil
}

func (e *MingEnv) AppEnv() string {
	return e.Get("APP_ENV")
}

func (e *MingEnv) IsExist(key string) bool {
	_, ok := e.maps[key]
	return ok
}

func (e *MingEnv) Get(key string) string {
	if v, ok := e.maps[key]; ok {
		return v
	}
	return ""
}

func (e *MingEnv) All() map[string]string {
	return e.maps
}
