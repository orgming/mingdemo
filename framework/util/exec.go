package util

import "os"

func GetExecDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return dir + "/"
	}
	return ""
}
