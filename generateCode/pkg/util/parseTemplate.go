package util

import (
	"text/template"
	"os"
)

func ParseTemplate(dstPath, srcPath string, data interface{}) {
	srcTemplate, err := template.ParseFiles(srcPath)
	if err != nil {
		panic(nil)
	}

	dstFileStream, err := os.Create(dstPath)
	if err != nil {
		panic(nil)
	}

	err = srcTemplate.Execute(dstFileStream, data)
	if err != nil {
		panic(nil)
	}
}