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
