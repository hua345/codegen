package util

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建目录和权限
func PathMkDir(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Println("Path:", path, " mkdir failed!", err)
		os.Exit(1)
	}
}

//创建多级目录和设置权限
func PathMkdirAll(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("Path:", path, " mkdir failed!", err)
		os.Exit(1)
	}
}

//删除目录
func PathRemove(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

//删除多级目录
func PathRemoveAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}
