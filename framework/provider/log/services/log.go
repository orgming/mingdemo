package services

import (
	"context"
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/provider/log/formatter"
	"io"
	pkgLog "log"
	"time"
)

// MingLog 日志通用实例
type MingLog struct {
	level      contract.LogLevel
	formatter  contract.Formatter
	ctxFielder contract.CtxFielder
	output     io.Writer
	c          framework.Container
}

// IsLevelEnable 判断这个级别日志是否可以打印
func (log *MingLog) IsLevelEnable(level contract.LogLevel) bool {
	return level <= log.level
}

func (log *MingLog) logf(level contract.LogLevel, ctx context.Context, msg string, fields map[string]any) error {
	if !log.IsLevelEnable(level) {
		return nil
	}

	fs := fields
	if log.ctxFielder != nil {
		if t := log.ctxFielder(ctx); t != nil {
			for k, v := range t {
				fs[k] = v
			}
		}
	}

	// 如果绑定了trace服务，获取trace信息
	if log.c.IsBind(contract.TraceKey) {
		tracer := log.c.MustMake(contract.TraceKey).(contract.Trace)
		tc := tracer.GetTrace(ctx)
		if tc != nil {
			maps := tracer.ToMap(tc)
			for k, v := range maps {
				fs[k] = v
			}
		}
	}

	// 将日志信息按照formatter序列化为字符串
	if log.formatter == nil {
		log.formatter = formatter.TextFormatter
	}
	ct, err := log.formatter(level, time.Now(), msg, fs)
	if err != nil {
		return err
	}

	// 如果是panic级别，则使用log进行panic
	if level == contract.PanicLevel {
		pkgLog.Panicln(string(ct))
		return nil
	}

	// 通过output进行输出
	log.output.Write(ct)
	log.output.Write([]byte("\r\n"))
	return nil
}

func (log *MingLog) Panic(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.PanicLevel, ctx, msg, fields)
}

func (log *MingLog) Fatal(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.FatalLevel, ctx, msg, fields)
}

func (log *MingLog) Error(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.ErrorLevel, ctx, msg, fields)
}

func (log *MingLog) Warn(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.WarnLevel, ctx, msg, fields)
}

func (log *MingLog) Info(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.InfoLevel, ctx, msg, fields)
}

func (log *MingLog) Debug(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.DebugLevel, ctx, msg, fields)
}

func (log *MingLog) Trace(ctx context.Context, msg string, fields map[string]any) {
	log.logf(contract.TraceLevel, ctx, msg, fields)
}

func (log *MingLog) SetLevel(level contract.LogLevel) {
	log.level = level
}

func (log *MingLog) SetCtxFielder(handler contract.CtxFielder) {
	log.ctxFielder = handler
}

func (log *MingLog) SetFormatter(formatter contract.Formatter) {
	log.formatter = formatter
}

func (log *MingLog) SetOutput(out io.Writer) {
	log.output = out
}
