package util

import (
	"os"
	"text/template"
)

func ParseTemplate(dstPath, srcPath string, data interface{}) {
	srcTemplate, err := template.ParseFiles(srcPath)
	if err != nil {
		panic(nil)
	}

	dstFileStream, err := os.Create(dstPath)
	defer dstFileStream.Close()
	if err != nil {
		panic(nil)
	}

	err = srcTemplate.Execute(dstFileStream, data)
	if err != nil {
		panic(nil)
	}
}
func ParseMethodTemplate(srcPath string, data interface{}) string {
	var tempFile string = "tempFile"
	ParseTemplate(tempFile, srcPath, data)
	content := ReadFileWithIoUtil(tempFile)
	defer os.Remove(tempFile)
	return content
}