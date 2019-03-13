package models

import (
	"fmt"
	"generateCode/config"
	"generateCode/pkg/util"
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
	fmt.Println("ProjectName:", projectInfoDto.ProjectName)
	fmt.Println("projectInfo:", projectInfoDto)

	projectInfoDto.CheckProjectPath()
	return projectInfoDto
}
func (projectInfoDto ProjectInfoDto) CheckProjectPath() {
	// 检测项目文件夹是否存在
	exist, err := util.PathExists(projectInfoDto.ProjectName)
	if err != nil {
		panic(err)
	}
	if exist {
		fmt.Printf("Project %s has Exist!\n", projectInfoDto.ProjectName)
		//os.Exit(2)
	} else {
		util.PathMkDir(projectInfoDto.ProjectName)
	}
	// 检测package目录是否存在
	javaPathExist, err := util.PathExists(projectInfoDto.JavaPath)
	if err != nil {
		panic(err)
	}
	if !javaPathExist {
		util.PathMkdirAll(projectInfoDto.JavaPath)
	}
	// 检测Util目录是否存在
	JavaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
	javaUtilPathExist, err := util.PathExists(JavaCodeUtilPath)
	if err != nil {
		panic(err)
	}
	if !javaUtilPathExist {
		util.PathMkdirAll(JavaCodeUtilPath)
	}
	// 检测Config目录是否存在
	JavaCodeConfigPath := path.Join(projectInfoDto.JavaPath, config.JavaConfigPath)
	javaConfigPathExist, err := util.PathExists(JavaCodeConfigPath)
	if err != nil {
		panic(err)
	}
	if !javaConfigPathExist {
		util.PathMkdirAll(JavaCodeConfigPath)
	}
	// 检测resource目录是否存在
	resourcePathExist, err := util.PathExists(projectInfoDto.ResourcePath)
	if err != nil {
		panic(err)
	}
	if !resourcePathExist {
		util.PathMkdirAll(projectInfoDto.ResourcePath)
	}
	_, err = util.CopyFile(path.Join(projectInfoDto.ProjectName, config.GitIgnoreFileName),
		path.Join(config.JavaTemplateInitPath, config.GitIgnoreFileName))
	if err != nil {
		fmt.Printf("Create Project dir %s failed!", )
	}
	initProjectData(projectInfoDto)
}

func initProjectData(projectInfoDto ProjectInfoDto) {
	fileMapDtoList := []FileMapDto{
		FileMapDto{path.Join(config.JavaTemplateInitPath, config.PomXmlFileName),
			path.Join(projectInfoDto.ProjectName, config.PomXmlFileName)},
		FileMapDto{path.Join(config.JavaTemplateInitPath, config.DotProjectFileName),
			path.Join(projectInfoDto.ProjectName, config.DotProjectFileName)},
		FileMapDto{path.Join(config.JavaTemplateCodePath, config.JavaApplicationFileName),
			path.Join(projectInfoDto.JavaPath, config.JavaApplicationFileName)},
	}
	// Java Util包
	JavaTemplateUtilPath := path.Join(config.JavaTemplateCodePath, config.JavaUtilPath)
	JavaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
	utilFileNameList := util.GetFilesName(JavaTemplateUtilPath)
	for _, value := range utilFileNameList {
		fileMapDtoList = append(fileMapDtoList,
			FileMapDto{path.Join(JavaTemplateUtilPath, value),
				path.Join(JavaCodeUtilPath, value)})
	}
	// Java Config包
	JavaTemplateConfigPath := path.Join(config.JavaTemplateCodePath, config.JavaConfigPath)
	JavaCodeConfigPath := path.Join(projectInfoDto.JavaPath, config.JavaConfigPath)
	configFileNameList := util.GetFilesName(JavaTemplateConfigPath)
	for _, value := range configFileNameList {
		fileMapDtoList = append(fileMapDtoList,
			FileMapDto{path.Join(JavaTemplateConfigPath, value),
				path.Join(JavaCodeConfigPath, value)})
	}
	// Resource文件
	resourceFileNameList := util.GetFilesName(config.JavaTemplateResourcePath)
	for _, value := range resourceFileNameList {
		fileMapDtoList = append(fileMapDtoList,
			FileMapDto{path.Join(config.JavaTemplateResourcePath, value),
				path.Join(projectInfoDto.ResourcePath, value)})
	}
	// 解析模板
	for _, value := range fileMapDtoList {
		util.ParseTemplate(value.TplDstPath, value.TplSrcPath, projectInfoDto)
	}
}
