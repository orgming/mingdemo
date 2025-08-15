package services

import "os"

type MingSingleLog struct {
	MingLog

	folder string
	file   string
	fd     *os.File
}

func NewMingSingleLog(params ...any) (any, error) {
	// TODO
	return nil, nil
}
