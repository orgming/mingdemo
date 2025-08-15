package services

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"os"
)

type MingConsoleLog struct {
	MingLog
}

func NewMingConsoleLog(params ...any) (any, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &MingConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
