package models

import (
	"encoding/json"
	"fmt"
	"codegen/pkg/util"
	"codegen/pkg/config"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

type SpringBootRestfulApiDto struct {
	HttpMethod            string         `json:"httpMethod"`
	BaseUrl               string         `json:"baseUrl"`
	RestfulUrl            string         `json:"restfulUrl"`
	MethodName            string         `json:"methodName"`
	MethodURL             string         `json:"methodURL"`
	ControllerName        string         `json:"controllerName"`
	RequestDTOName        string         `json:"requestDTOName"`
	ResponseDTOName       string                   `json:"responseDTOName"`
	VarResponseDTOName    string                   `json:"varResponseDTOName"`
	ImportRequestDTOPath  string                   `json:"importRequestDTOPath"`
	ImportResponseDTOPath string                   `json:"importResponseDTOPath"`
	ControllerURL         string                   `json:"controllerURL"`
	Description           string                   `json:"description"`
	NowDate               string                   `json:"nowDate"`
	Author                string                   `json:"author"`
	ProjectInfo           SpringBootProjectInfoDto `json:"projectInfo"`
}

func (restfulApiDto SpringBootRestfulApiDto) Init() SpringBootRestfulApiDto {
	restfulApiDto.NowDate = time.Now().Format(config.NowTimeFormat)
	restfulApiDto.Author = config.ServerConfig.AuthorName
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
		restfulApiDto.HttpMethod = config.HttpMethodMapping[config.ServerConfig.DefaultHttpMethod]
	}
	restfulApiDto.MethodName = util.StrFirstToLower(restfulApiDto.MethodName)
	// Description
	if len(restfulApiDto.Description) == 0 {
		restfulApiDto.Description = restfulApiDto.MethodName
	}
	// DTO
	restfulApiDto.RequestDTOName = util.StrFirstToUpper(restfulApiDto.MethodName) + config.ImportRequestDto
	restfulApiDto.ResponseDTOName = util.StrFirstToUpper(restfulApiDto.MethodName) + config.ImportResponseDto
	restfulApiDto.VarResponseDTOName = restfulApiDto.MethodName + config.ImportResponseDto
	restfulApiDto.ImportRequestDTOPath = config.ImportPrefix + restfulApiDto.ProjectInfo.PackageName + "." +
		config.ImportDtoRequestPath + "." + restfulApiDto.RequestDTOName + ";"
	restfulApiDto.ImportResponseDTOPath = config.ImportPrefix + restfulApiDto.ProjectInfo.PackageName + "." +
		config.ImportDtoResponsePath + "." + restfulApiDto.ResponseDTOName + ";"
	return restfulApiDto
}
func (restfulApiDto SpringBootRestfulApiDto) GenerateCode() {
	data, _ := json.MarshalIndent(restfulApiDto, "", "    ")
	fmt.Printf("%s\n", data)
	// 检测项目文件夹是否存在
	existProject, err := util.PathExists(restfulApiDto.ProjectInfo.ProjectName)
	if err != nil {
		panic(err)
	}
	if !existProject {
		fmt.Printf("Project %s Not Found")
		os.Exit(2)
	}
	controllerExist := checkControllerExist(restfulApiDto)
	if controllerExist {
		restfulAddMethod(restfulApiDto)
	} else {
		restfulApiNew(restfulApiDto)
	}
}
func checkControllerExist(restfulApiDto SpringBootRestfulApiDto) (controllerExist bool) {
	// 检测controller目录是否存在
	codeControllerPath := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaControllerPath)
	util.CheckDirAndMkdir(codeControllerPath)
	// restful controller文件名
	controllerFilePath := path.Join(codeControllerPath, restfulApiDto.ControllerName+config.JavaTemplateControllerFileName)
	controllerExist, err := util.PathExists(controllerFilePath)
	if err != nil {
		panic(err)
	}
	if !controllerExist {
		return false
	} else {
		return true
	}
}

/**
 * 检查Controller是否存在相同方法
 */
func checkSameMethod(controllerContent string, restfulApiDto SpringBootRestfulApiDto) {
	// 获取当前Controller所有方法
	contentReg := regexp.MustCompile(`public.*\(`)
	mappingReg := regexp.MustCompile(`@.*` + config.SpringMapping + `.*\)`)
	// 方法正则
	methodRegList := contentReg.FindAllString(controllerContent, -1)
	// URL和请求方式正则
	mappingRegList := mappingReg.FindAllString(controllerContent, -1)
	mappingRegList = mappingRegList[1:]
	if len(methodRegList) != len(mappingRegList) {
		fmt.Printf("Controller: %s的方法和请求方式与请求路径不能一一对应!\n", restfulApiDto.ControllerName)
		os.Exit(6)
	}
	var methodMappingInfoSlice []MethodMappingInfo
	for index, value := range methodRegList {
		var methodMappingInfo MethodMappingInfo
		// 截取方法名称
		methodSplitList := strings.Split(value, " ")
		if len(methodSplitList) != 3 {
			fmt.Printf("解析(%s)获取方法名称失败!", value)
			os.Exit(6)
		}
		methodName := strings.Replace(methodSplitList[len(methodSplitList)-1], "(", "", 1)
		methodMappingInfo.MethodName = methodName
		// 截取请求方式和URL
		mappingValue := mappingRegList[index]
		mappingValueList := strings.Split(mappingValue, config.SpringMapping)
		if len(mappingValueList) != 2 {
			fmt.Printf("解析(%s)获取请求方式和URL失败!", mappingValue)
			os.Exit(6)
		}
		// 截取请求方式
		httpMethod := strings.Replace(mappingValueList[0], "@", "", 1)
		methodMappingInfo.HttpMethod = httpMethod
		//截取请求URL
		urlPathList := strings.Split(mappingValueList[1], "=")
		if len(urlPathList) != 2 {
			fmt.Printf("解析(%s)获取方法URL失败!", mappingValue)
			os.Exit(6)
		}
		urlPath := strings.TrimSpace(strings.Replace(urlPathList[1], ")", "", 1))
		urlPath = strings.Replace(urlPath, "\"", "", 10)
		methodMappingInfo.UrlPath = urlPath
		methodMappingInfoSlice = append(methodMappingInfoSlice, methodMappingInfo)
	}
	data, err := json.MarshalIndent(methodMappingInfoSlice, "", "    ")
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("接口列表: %s\n", data)
	// 检查是否已经存在同一名称的方法
	for _, item := range methodMappingInfoSlice {
		if strings.ToLower(item.MethodName) == strings.ToLower(restfulApiDto.MethodName) {
			fmt.Printf("Controller %s 已经存在方法 %s\n", restfulApiDto.ControllerName, restfulApiDto.MethodName)
			os.Exit(5)
		}
	}
	// 检查是否存在同一请求路径和请求方法
	for _, item := range methodMappingInfoSlice {
		if strings.ToLower(item.UrlPath) == strings.ToLower(restfulApiDto.MethodURL) &&
			strings.ToLower(item.HttpMethod+config.SpringMapping) == strings.ToLower(restfulApiDto.HttpMethod) {
			fmt.Printf("Controller(%s)已经存在同一请求路径(%s)和请求方法(%s)\n",
				restfulApiDto.ControllerName, restfulApiDto.MethodName, item.HttpMethod)
			os.Exit(5)
		}
	}
}

/**
 * 添加在同一个Controller时的方法
 */
func restfulAddMethod(restfulApiDto SpringBootRestfulApiDto) {
	// 检测controller目录是否存在
	codeControllerDir := path.Join(restfulApiDto.ProjectInfo.JavaPath, config.JavaControllerPath)
	util.CheckDirAndMkdir(codeControllerDir)
	// restful controller文件名
	controllerFilePath := path.Join(codeControllerDir, restfulApiDto.ControllerName+config.JavaTemplateControllerFileName)
	controllerContent := util.ReadFileWithIoUtil(controllerFilePath)
	checkSameMethod(controllerContent, restfulApiDto)
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

	var fileMapDtoList []FileMapDto
	// request dto Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.ImportRequestDto+config.JavaSuffixName),
			path.Join(codeDtoRequestPath, restfulApiDto.RequestDTOName+config.JavaSuffixName)})
	// response dto Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.ImportResponseDto+config.JavaSuffixName),
			path.Join(codeDtoResponsePath, restfulApiDto.ResponseDTOName+config.JavaSuffixName)})
	// 解析模板
	for _, value := range fileMapDtoList {
		util.ParseTemplate(value.TplDstPath, value.TplSrcPath, restfulApiDto)
	}

	// 添加方法
	methodCode := util.ParseMethodTemplate(path.Join(config.JavaTemplateCodePath, config.JavaTemplateMethodControllerFileName), restfulApiDto)
	dstControllerFileName := path.Join(codeControllerDir, restfulApiDto.ControllerName+config.JavaTemplateControllerFileName)
	appendMethod(methodCode, dstControllerFileName, restfulApiDto, false)

	serviceMethodCode := util.ParseMethodTemplate(path.Join(config.JavaTemplateCodePath, config.JavaTemplateMethodServiceFileName), restfulApiDto)
	dstServicePath := path.Join(codeServicePath, restfulApiDto.ControllerName+config.JavaTemplateServiceFileName)
	appendMethod(serviceMethodCode, dstServicePath, restfulApiDto, true)

	serviceImplMethodCode := util.ParseMethodTemplate(path.Join(config.JavaTemplateCodePath, config.JavaTemplateMethodServiceImplFileName), restfulApiDto)
	dstServiceImplPath := path.Join(codeServiceImplPath, restfulApiDto.ControllerName+config.JavaTemplateServiceImplFileName)
	appendMethod(serviceImplMethodCode, dstServiceImplPath, restfulApiDto, false)
}
func appendMethod(methodCode, dstFilePath string, restfulApiDto SpringBootRestfulApiDto, javaInterface bool) {
	srcContent := util.ReadFileWithIoUtil(dstFilePath)
	// }加换行符
	contentReg := regexp.MustCompile(`}[\n|\r\n]`)
	srcContentSlice := contentReg.Split(srcContent, -1)

	resultControllerContent := srcContentSlice[0 : len(srcContentSlice)-1]
	resultControllerContent = append(resultControllerContent, methodCode+config.RowLimiter)
	var resultContent string
	for _, value := range resultControllerContent {
		if len(strings.TrimSpace(value)) >= 3 {
			if javaInterface {
				resultContent = resultContent + value
			} else {
				resultContent = resultContent + value + "}" + config.RowLimiter
			}
		}
	}
	if javaInterface {
		resultContent = resultContent + "}" + config.RowLimiter
	}
	oldRequest := config.ImportRequestDto + ";"
	resultContent = strings.Replace(resultContent, oldRequest, oldRequest+
		config.RowLimiter+restfulApiDto.ImportRequestDTOPath, 1)
	oldResponse := config.ImportResponseDto + ";"
	resultContent = strings.Replace(resultContent, oldResponse, oldResponse+
		config.RowLimiter+restfulApiDto.ImportResponseDTOPath, 1)

	importAnnotation := config.ImportSpringAnnotation + restfulApiDto.HttpMethod + ";"
	if !strings.Contains(resultContent, importAnnotation) {
		defaultAnnotation := config.HttpMethodMapping[config.ServerConfig.DefaultHttpMethod] + ";"
		resultContent = strings.Replace(resultContent, defaultAnnotation, defaultAnnotation+
			config.RowLimiter+importAnnotation, 1)
	}

	util.WriteFileWithIoUtil(dstFilePath, resultContent)
}

// 创建全新的接口文件
func restfulApiNew(restfulApiDto SpringBootRestfulApiDto) {
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
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaTemplateControllerFileName),
			path.Join(codeControllerPath, restfulApiDto.ControllerName+config.JavaTemplateControllerFileName)})
	// request dto Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.ImportRequestDto+config.JavaSuffixName),
			path.Join(codeDtoRequestPath, restfulApiDto.RequestDTOName+config.JavaSuffixName)})
	// response dto Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.ImportResponseDto+config.JavaSuffixName),
			path.Join(codeDtoResponsePath, restfulApiDto.ResponseDTOName+config.JavaSuffixName)})

	// Service Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaTemplateServiceFileName),
			path.Join(codeServicePath, restfulApiDto.ControllerName+config.JavaTemplateServiceFileName)})
	// ServiceImpl Code
	fileMapDtoList = append(fileMapDtoList,
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaTemplateServiceImplFileName),
			path.Join(codeServiceImplPath, restfulApiDto.ControllerName+config.JavaTemplateServiceImplFileName)})
	fmt.Println(fileMapDtoList)

	// 解析模板
	for _, value := range fileMapDtoList {
		util.ParseTemplate(value.TplDstPath, value.TplSrcPath, restfulApiDto)
	}
}
