package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetCurrentPath 获取当前绝对路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, string(os.PathSeparator))
	if i < 0 {
		return "", errors.New("error: can't find " + string(os.PathSeparator))
	}
	return path[0 : i+1], nil
}

// GetWdPath ...
func GetWdPath() (string, error) {
	file, err := os.Getwd()
	if err != nil {
		return "", nil
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return path, err
}

//GetFilePath ...
func GetFilePath(file string) (string, error) {
	basePath, err := GetCurrentPath()
	pathFile := basePath + file
	_, err = os.Stat(pathFile)
	// == nil 表示文件存在
	if err == nil {
		return pathFile, err
	}
	if os.IsNotExist(err) {
		// 不存在就从当前目前下找
		path, err := GetWdPath()
		if err != nil {
			return "", nil
		}
		return path + string(os.PathSeparator) + file, nil
	}
	// 再不存在，就报错
	return "", err
}
