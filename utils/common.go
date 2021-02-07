package utils

import (
	"os"
	"path/filepath"
	"time"
)

func CheckERR(err error) {
	if err != nil {
		panic(err)
	}
}

func PointNow() string {
	return time.Now().Format("2006.01.02.15.04.05")
}

func GetDirNow() string {
	// 获取程序所在目录
	dirNow, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dirNow
}
