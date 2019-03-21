package models

import (
	"fmt"
	"generateCode/config"
	"generateCode/pkg/util"
	"strings"
)

type RestfulApiDto struct {
	HttpMethod     string         `json:httpMethod`
	BaseUrl        string         `json:baseUrl`
	RestfulUrl     string         `json:restfulUrl`
	MethodName     string         `json:methodName`
	ControllerName string         `json:controllerName`
	Description    string         `json:description`
	ProjectInfo    ProjectInfoDto `json:projectInfo`
}

func (restfulApiDto RestfulApiDto) Init() RestfulApiDto {
	fmt.Println(restfulApiDto)
	//Http Method 处理
	httpMethod := strings.ToUpper(restfulApiDto.HttpMethod)
	restfulApiDto.HttpMethod = config.HttpMethodMapping[httpMethod]
	if len(restfulApiDto.HttpMethod) == 0 {
		restfulApiDto.HttpMethod = config.HttpMethodMapping["GET"]
	}
	restfulApiDto.MethodName = util.StrFirstToLower(restfulApiDto.MethodName)
	// 根据输入的RestfulUrl判断生成的Controller

	return restfulApiDto
}
