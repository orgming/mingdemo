package log

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/provider/log/formatter"
	"github.com/orgming/ming/framework/provider/log/services"
	"io"
	"strings"
)

type MingLogProvider struct {
	framework.ServiceProvider

	Driver string // Driver

	// 日志级别
	Level contract.LogLevel
	// 日志输出格式方法
	Formatter contract.Formatter
	// 日志context上下文信息获取函数
	CtxFielder contract.CtxFielder
	// 日志输出信息
	Output io.Writer
}

func (l *MingLogProvider) Register(c framework.Container) framework.NewInstance {
	if l.Driver == "" {
		tcs, err := c.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用console
			return services.NewMingConsoleLog
		}

		cs := tcs.(contract.Config)
		l.Driver = strings.ToLower(cs.GetString("log.Driver"))
	}

	// 根据driver的配置项确定
	switch l.Driver {
	case "single":
		return services.NewMingSingleLog
	case "rotate":
		return services.NewMingRotateLog
	case "console":
		return services.NewMingConsoleLog
	case "custom":
		return services.NewMingCustomLog
	default:
		return services.NewMingConsoleLog
	}
}

func (l *MingLogProvider) Boot(c framework.Container) error {
	return nil
}

func (l *MingLogProvider) IsDefer() bool {
	return false
}

func (l *MingLogProvider) Params(c framework.Container) []any {
	// 获取configService
	configService := c.MustMake(contract.ConfigKey).(contract.Config)

	// 设置参数formatter
	if l.Formatter == nil {
		l.Formatter = formatter.TextFormatter
		if configService.IsExist("log.formatter") {
			v := configService.GetString("log.formatter")
			if v == "json" {
				l.Formatter = formatter.JsonFormatter
			} else if v == "text" {
				l.Formatter = formatter.TextFormatter
			}
		}
	}

	if l.Level == contract.UnknownLevel {
		l.Level = contract.InfoLevel
		if configService.IsExist("log.level") {
			l.Level = logLevel(configService.GetString("log.level"))
		}
	}

	// 定义5个参数
	return []any{c, l.Level, l.CtxFielder, l.Formatter, l.Output}
}

func (l *MingLogProvider) Name() string {
	return contract.LogKey
}

func logLevel(config string) contract.LogLevel {
	switch strings.ToLower(config) {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	}
	return contract.UnknownLevel
}
