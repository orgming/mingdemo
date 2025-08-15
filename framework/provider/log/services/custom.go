package services

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"io"
)

type MingCustomLog struct {
	MingLog
}

func NewMingCustomLog(params ...any) (any, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &MingCustomLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c
	return log, nil
}
