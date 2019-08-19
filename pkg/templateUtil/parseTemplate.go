package templateUtil

import (
	"codegen/asset"
	"codegen/pkg/fileUtil"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

const publicDir = "public"

//获取文件夹下子文件名称列表
// public/js/a.js, public/js/b.js, public/js/c/c.js
// asset.AssetDir(public/js) => a.js b.js
// GetTemplateFilesName(js) => a.js b.js
func GetTemplateFilesName(assertPath string) []string {
	var fileNameList []string
	assertFileNameList, err := asset.AssetDir(path.Join(publicDir, assertPath))
	if err != nil {
		panic(err)
	}
	for _, fileItem := range assertFileNameList {
		fileInfo, _ := asset.AssetInfo(path.Join(publicDir, assertPath, fileItem))
		if nil == fileInfo || fileInfo.IsDir() {
			continue
		} else {
			fileNameList = append(fileNameList, fileItem)
		}
	}
	return fileNameList
}

//获取文件夹下子文件夹名称列表
// public/js/a.js, public/js/b.js, public/js/c/c.js
// GetTemplateSubDirList(js) => c
func GetTemplateSubDirList(assertPath string) []string {
	var subDirList []string
	assertFileNameList, err := asset.AssetDir(path.Join(publicDir, assertPath))
	if err != nil {
		panic(err)
	}
	for _, fileItem := range assertFileNameList {
		fileInfo, _ := asset.AssetInfo(path.Join(publicDir, assertPath, fileItem))
		if nil == fileInfo || fileInfo.IsDir() {
			subDirList = append(subDirList, fileItem)
		} else {
			continue
		}
	}
	return subDirList
}
func CopyTemplateFile(dstPath, assertPath string) {
	assertFile, err := asset.Asset(path.Join(publicDir, assertPath))
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(dstPath, assertFile, 0644)
	if err != nil {
		panic(err)
	}
}
func ParseTemplate(dstPath, srcPath string, data interface{}) {
	log.Println("srcPath: " + srcPath)
	log.Println("dstPath: " + dstPath)
	//**********************go-bindata begin ***********************//
	templateContent, err := asset.Asset(path.Join(publicDir, srcPath))
	if err != nil {
		log.Println(path.Join(publicDir, srcPath) + " Template File not found.")
		os.Exit(-1)
	}
	srcTemplate, err := template.New(srcPath).Parse(string(templateContent))
	if err != nil {
		panic(nil)
	}
	//**********************go-bindata end ***********************//
	//srcTemplate, err := template.ParseFiles(templateFile)
	//if err != nil {
	//	panic(nil)
	//}

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
	content := fileUtil.ReadFileWithIoUtil(tempFile)
	defer os.Remove(tempFile)
	return content
}
