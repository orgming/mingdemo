package util

import (
	"fmt"
	"os"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	type args struct {
		arr [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				arr: [][]string{
					{"te", "test", "sdf"},
					{"te11232", "test123123", "1232123"},
					{"test1xxx1232", "test123123", "1232123"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrettyPrint(tt.args.arr)
		})
	}
}

func TestCheckProcessExist(t *testing.T) {
	pid := os.Getpid()
	if CheckProcessExist(pid) {
		fmt.Println("当前进程的pid为： ", pid)
	}
}

func TestGetExecDirectory(t *testing.T) {
	fmt.Println(GetExecDirectory())
}
