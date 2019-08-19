package fileUtil

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

//文件打开模式：
//
//const (
//	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
//	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
//	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
//	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
//	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
//	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
//	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
//	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
//)
//权限控制：
//
//r ——> 004
//w ——> 002
//x ——> 001
// 拷贝文件，返回拷贝字节数
func CopyFile(dstPath, srcPath string) (nBytes int64, err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY, 0644)
	//操作完毕，关闭文件
	defer srcFile.Close()
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// 获取文件后缀
func GetFileSuffix(fileName string) string {
	var fileSuffix string
	fileSuffix = path.Ext(fileName) //获取文件后缀
	return fileSuffix
}

//获取文件夹文件列表
func GetFilesName(path string) []string {
	var names []string
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			names = append(names, file.Name())
		}
	}
	return names
}

/**
 * 检查目录是否存在，如果不存在创建目录
 */
func CheckDirAndMkdir(path string) {
	pathExist, err := PathExists(path)
	if err != nil {
		panic(err)
	}
	if !pathExist {
		PathMkdirAll(path)
	}
}

//获取文件夹下子文件夹列表
func GetSubDirList(path string) []string {
	var subDirList []string
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			subDirList = append(subDirList, file.Name())
		} else {
			continue
		}
	}
	return subDirList
}

// 读取文件
func ReadFileWithIoUtil(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}

// 写入文件
func WriteFileWithIoUtil(name, content string) {
	data := []byte(content)
	err := ioutil.WriteFile(name, data, 0644)
	if err != nil {
		panic(err)
	}
}
