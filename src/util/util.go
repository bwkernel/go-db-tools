package util

import (
	"os"
	"path/filepath"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if !os.IsExist(err) {
			return false
		}
	}

	if !IsFile(path) {
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
} // 判断所给路径是否为文件

func IsFile(path string) bool {
	return !IsDir(path)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径
	filepath.Dir(os.Args[0])
	if err != nil {
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
