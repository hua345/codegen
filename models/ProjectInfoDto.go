package models

import (
	"encoding/json"
	"fmt"
	"github.com/hua345/generateCode/config"
	"github.com/hua345/generateCode/pkg/util"
	"os"
	"path"
	"strings"
)

type ProjectInfoDto struct {
	GroupId      string `json:"groupdId"`
	ArtifactId   string `json:"artifactId"`
	PackageName  string `json:"packageName"`
	JavaPath     string `json:"javaPath"`
	ResourcePath string `json:"resourcePath"`
	ProjectName  string `json:"projectName"`
}

// 初始化项目数据
func (projectInfoDto ProjectInfoDto) Init() ProjectInfoDto {
	// 包路径名称小写
	projectInfoDto.GroupId = strings.ToLower(projectInfoDto.GroupId)
	projectInfoDto.ArtifactId = strings.ToLower(projectInfoDto.ArtifactId)

	artifactList := strings.Split(projectInfoDto.ArtifactId, ".")
	projectInfoDto.ProjectName = ""
	for _, value := range artifactList {
		if len(projectInfoDto.ProjectName) >= 1 {
			projectInfoDto.ProjectName = projectInfoDto.ProjectName + "-" + value
		} else {
			projectInfoDto.ProjectName = value
		}
	}
	// 连接符'-'替换
	projectInfoDto.GroupId = strings.Replace(projectInfoDto.GroupId, "-", "", -1)
	projectInfoDto.ArtifactId = strings.Replace(projectInfoDto.ArtifactId, "-", "", -1)
	projectInfoDto.PackageName = projectInfoDto.GroupId +
		"." + projectInfoDto.ArtifactId

	// package目录
	projectInfoDto.JavaPath = path.Join(projectInfoDto.ProjectName, config.JavaPath)
	projectInfoDto.ResourcePath = path.Join(projectInfoDto.ProjectName, config.JavaResourcePath)
	packageNameList := strings.Split(projectInfoDto.PackageName, ".")
	for _, value := range packageNameList {
		projectInfoDto.JavaPath = path.Join(projectInfoDto.JavaPath, value)
	}
	return projectInfoDto
}
func (projectInfoDto ProjectInfoDto) InitProject() {
	data, _ := json.MarshalIndent(projectInfoDto, "", "    ")
	fmt.Printf("%s\n", data)
	// 检测项目文件夹是否存在
	exist, err := util.PathExists(projectInfoDto.ProjectName)
	if err != nil {
		panic(err)
	}
	if exist {
		fmt.Printf("Project %s has Exist!\n", projectInfoDto.ProjectName)
		os.Exit(2)
	} else {
		util.PathMkDir(projectInfoDto.ProjectName)
	}
	// 检测package目录是否存在
	util.CheckDirAndMkdir(projectInfoDto.JavaPath)
	// 检测Util目录是否存在
	JavaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
	util.CheckDirAndMkdir(JavaCodeUtilPath)
	// 检测Config目录是否存在
	JavaCodeConfigPath := path.Join(projectInfoDto.JavaPath, config.JavaConfigPath)
	util.CheckDirAndMkdir(JavaCodeConfigPath)
	// 检测Common目录是否存在
	JavaCodeCommonPath := path.Join(projectInfoDto.JavaPath, config.JavaCommonPath)
	util.CheckDirAndMkdir(JavaCodeCommonPath)
	// 检测resource目录是否存在
	util.CheckDirAndMkdir(projectInfoDto.ResourcePath)
	// 检测mybatis目录是否存在
	mybatisPath := path.Join(projectInfoDto.ResourcePath, config.MybatisPath)
	util.CheckDirAndMkdir(mybatisPath)
	// Copy .Gitignore
	_, err = util.CopyFile(path.Join(projectInfoDto.ProjectName, config.GitIgnoreFileName),
		path.Join(config.JavaTemplateInitPath, config.GitIgnoreFileName))
	if err != nil {
		fmt.Printf("Copy %s failed!", config.GitIgnoreFileName)
	}
	// 解析模板
	initProjectData(projectInfoDto)
}

/**
 * 添加common文件模板映射关系
 */
func appendCommonTemplateList(templateConfigPath, codeConfigPath string, fileMapDtoList []FileMapDto) []FileMapDto {
	subDirList := util.GetSubDirList(templateConfigPath)
	for _, value := range subDirList {
		templateSubPath := path.Join(templateConfigPath, value)
		codeSubDir := path.Join(codeConfigPath, value)
		util.CheckDirAndMkdir(codeSubDir)
		fileMapDtoList = appendTemplateList(templateSubPath, codeSubDir, fileMapDtoList)
	}
	return fileMapDtoList
}

/**
 * 添加Mybatis文件模板映射关系
 */
func appendMybatisTemplateList(mybatisTemplatePath, mybatisPath string, fileMapDtoList []FileMapDto) []FileMapDto {
	mybatisFileNameList := util.GetFilesName(mybatisTemplatePath)
	for _, value := range mybatisFileNameList {
		if util.GetFileSuffix(value) == ".jar" {
			// 如果是jar文件直接拷贝
			_, err := util.CopyFile(path.Join(mybatisPath, value),
				path.Join(mybatisTemplatePath, value))
			if err != nil {
				fmt.Printf("Copy %s failed!", value)
			}
		} else {
			fileMapDtoList = append(fileMapDtoList,
				FileMapDto{path.Join(mybatisTemplatePath, value),
					path.Join(mybatisPath, value)})
		}
	}
	return fileMapDtoList
}

/**
 * 添加文件模板映射关系
 */
func appendTemplateList(templatePath, codePath string, fileMapDtoList []FileMapDto) []FileMapDto {
	templateFileNameList := util.GetFilesName(templatePath)
	for _, value := range templateFileNameList {
		fileMapDtoList = append(fileMapDtoList,
			FileMapDto{path.Join(templatePath, value),
				path.Join(codePath, value)})
	}
	return fileMapDtoList
}
func initProjectData(projectInfoDto ProjectInfoDto) {
	fileMapDtoList := []FileMapDto{
		FileMapDto{path.Join(config.JavaTemplateInitPath, config.PomXmlFileName),
			path.Join(projectInfoDto.ProjectName, config.PomXmlFileName)},
		FileMapDto{path.Join(config.JavaTemplateInitPath, config.DotProjectFileName),
			path.Join(projectInfoDto.ProjectName, config.DotProjectFileName)},
		FileMapDto{path.Join(config.JavaTemplateInitCodePath, config.JavaApplicationFileName),
			path.Join(projectInfoDto.JavaPath, config.JavaApplicationFileName)},
	}
	// Java Util包
	javaTemplateUtilPath := path.Join(config.JavaTemplateInitCodePath, config.JavaUtilPath)
	javaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
	fileMapDtoList = appendTemplateList(javaTemplateUtilPath, javaCodeUtilPath, fileMapDtoList)
	// Java Config包
	javaTemplateConfigPath := path.Join(config.JavaTemplateInitCodePath, config.JavaConfigPath)
	javaCodeConfigPath := path.Join(projectInfoDto.JavaPath, config.JavaConfigPath)
	fileMapDtoList = appendTemplateList(javaTemplateConfigPath, javaCodeConfigPath, fileMapDtoList)
	// Config subDir文件
	fileMapDtoList = appendCommonTemplateList(javaTemplateConfigPath, javaCodeConfigPath, fileMapDtoList)
	// Java Common包
	javaTemplateCommonPath := path.Join(config.JavaTemplateInitCodePath, config.JavaCommonPath)
	javaCodeCommonPath := path.Join(projectInfoDto.JavaPath, config.JavaCommonPath)
	fileMapDtoList = appendTemplateList(javaTemplateCommonPath, javaCodeCommonPath, fileMapDtoList)
	// Resource文件
	fileMapDtoList = appendTemplateList(config.JavaTemplateResourcePath, projectInfoDto.ResourcePath, fileMapDtoList)
	// Mybatis
	mybatisTemplatePath := path.Join(config.JavaTemplateInitPath, config.MybatisPath)
	mybatisPath := path.Join(projectInfoDto.ResourcePath, config.MybatisPath)
	fileMapDtoList = appendMybatisTemplateList(mybatisTemplatePath, mybatisPath, fileMapDtoList)

	// 解析模板
	for _, value := range fileMapDtoList {
		util.ParseTemplate(value.TplDstPath, value.TplSrcPath, projectInfoDto)
	}
}
