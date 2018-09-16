// jsonCode project jsonCode.go
package jsonCode

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func ReadJsonData() []ServiceInfo {
	
	var infoList []ServiceInfo
	jsonFileList := getFilesName(path.Join(JSONPATH))
	for _, item := range jsonFileList {
		var info ServiceInfo
		serviceName := getServiceName(item)
		fmt.Println(info)
		fmt.Println(serviceName)
	}
	return infoList
}
func getServiceName(fileName string) string {
	ArtifactList := strings.Split(fileName, "_")
	if len(ArtifactList) < 3 {
		fmt.Println("JSON File Name:", fileName)
		fmt.Println("For example: demo_user_xxx.json")
		os.Exit(3)
	}

	serviceName := ArtifactList[0] + "." + ArtifactList[1] + "." + ArtifactList[2]
	return strings.ToLower(serviceName)
}

func getFilesName(path string) []string {
	var filenameList []string
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			filenameList = append(filenameList, file.Name())
		}
	}
	return filenameList
}
