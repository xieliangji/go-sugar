package util

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallDir - 安装路径
func InstallDir() string {
	p, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Println(err)
		return ""
	}

	absPath, err := filepath.Abs(p)
	if err != nil {
		log.Println(err)
		return ""
	}

	index := strings.LastIndex(absPath, string(os.PathSeparator))
	return absPath[:index]
}

// WorkDir - 工作路径
func WorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return ""
	}
	return dir
}
