package models

import (
	"encoding/json"
	"fmt"
	"generateCode/config"
	"generateCode/pkg/util"
	"os"
	"path"
	"strings"
)

type RestfulApiDto struct {
	HttpMethod     string         `json:httpMethod`
	BaseUrl        string         `json:baseUrl`
	RestfulUrl     string         `json:restfulUrl`
	MethodName     string         `json:methodName`
	MethodURL      string         `json:methodURL`
	ControllerName string         `json:controllerName`
	DTOName        string         `json:dtoName`
	ControllerURL  string         `json:controllerURL`
	Description    string         `json:description`
	ProjectInfo    ProjectInfoDto `json:projectInfo`
}

func (restfulApiDto RestfulApiDto) Init() RestfulApiDto {
	// 根据输入的RestfulUrl判断生成的Controller
	urlStrList := util.HandleRestfulURL(restfulApiDto.RestfulUrl)
	if len(urlStrList) == 0 {
		fmt.Println("URL路径： " + restfulApiDto.RestfulUrl + "不符合规范")
		os.Exit(3)
	}
	// URL处理
	restfulApiDto.ControllerName = util.StrFirstToUpper(urlStrList[0])
	restfulApiDto.ControllerURL = util.AppendURL(restfulApiDto.BaseUrl, urlStrList[0])
	methodUrlList := urlStrList[1:]
	if len(methodUrlList) == 0 {
		restfulApiDto.MethodURL = "/"
	} else {
		restfulApiDto.MethodURL = ""
		for _, value := range methodUrlList {
			restfulApiDto.MethodURL = restfulApiDto.MethodURL + "/" + value
		}
	}
	//Http Method 处理
	if len(restfulApiDto.HttpMethod) == 0 {
		restfulApiDto.HttpMethod = urlStrList[0]
	}
	httpMethod := strings.ToUpper(restfulApiDto.HttpMethod)
	restfulApiDto.HttpMethod = config.HttpMethodMapping[httpMethod]
	if len(restfulApiDto.HttpMethod) == 0 {
		restfulApiDto.HttpMethod = config.HttpMethodMapping[config.DefaultHttpMethod]
	}
	restfulApiDto.MethodName = util.StrFirstToLower(restfulApiDto.MethodName)
	restfulApiDto.DTOName = util.StrFirstToUpper(restfulApiDto.MethodName)
	// Description
	if len(restfulApiDto.Description) == 0 {
		restfulApiDto.Description = restfulApiDto.MethodName
	}
	return restfulApiDto
}
func (restfulApiDto RestfulApiDto) GenerateCode() {
	data,_ := json.Marshal(restfulApiDto)
	fmt.Printf("%s\n", data)
	// 检测项目文件夹是否存在
	existProject, err := util.PathExists(restfulApiDto.ProjectInfo.ProjectName)
	if err != nil {
		panic(err)
	}
	if !existProject {
		fmt.Printf("Project %s Not Found", )
		os.Exit(2)
	}
	// 检测controller目录是否存在
	codeControllerPath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaControllerPath)
	util.CheckDirAndMkdir(codeControllerPath)
	// 检测dto目录是否存在
	codeDtoPath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaDtoPath)
	util.CheckDirAndMkdir(codeDtoPath)
	// 检测dto request目录是否存在
	codeDtoRequestPath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaDtoRequestPath)
	util.CheckDirAndMkdir(codeDtoRequestPath)
	// 检测dto response目录是否存在
	codeDtoResponsePath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaDtoResponsePath)
	util.CheckDirAndMkdir(codeDtoResponsePath)
	// 检测service目录是否存在
	codeServicePath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaServicePath)
	util.CheckDirAndMkdir(codeServicePath)
	// 检测serviceImpl目录是否存在
	codeServiceImplPath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaServiceImplPath)
	util.CheckDirAndMkdir(codeServiceImplPath)
	//
	var fileMapDtoList []FileMapDto

	// Controller Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaControllerFileName),
			path.Join(codeControllerPath, restfulApiDto.ControllerName + config.JavaControllerFileName)})
	// request dto Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaRequestDtoFileName),
			path.Join(codeDtoRequestPath, restfulApiDto.DTOName + config.JavaRequestDtoFileName)})
	// response dto Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaResponseDtoFileName),
			path.Join(codeDtoResponsePath, restfulApiDto.DTOName + config.JavaResponseDtoFileName)})

	// Service Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaServiceFileName),
			path.Join(codeServicePath, restfulApiDto.ControllerName + config.JavaServiceFileName)})
	// ServiceImpl Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaServiceImplFileName),
			path.Join(codeServiceImplPath, restfulApiDto.ControllerName + config.JavaServiceImplFileName)})
	fmt.Println(fileMapDtoList)
	// 解析模板
	for _, value := range fileMapDtoList {
		util.ParseTemplate(value.TplDstPath, value.TplSrcPath, restfulApiDto)
	}
}
