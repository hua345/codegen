package models

import (
	"github.com/hua345/codegen/pkg/config"
	"github.com/hua345/codegen/pkg/fileUtil"
	"github.com/hua345/codegen/pkg/templateUtil"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type SpringBootProjectInfoDto struct {
	GroupId               string          `json:"groupdId"`
	ArtifactId            string          `json:"artifactId"`
	PackageName           string          `json:"packageName"`
	JavaPath              string          `json:"javaPath"`
	JavaTestPath          string          `json:"javaTestPath"`
	ResourcePath          string          `json:"resourcePath"`
	ProjectName           string          `json:"projectName"`
	NowDate               string          `json:"nowDate"`
	Author                string          `json:"author"`
	SupportMaven          bool            `yaml:"supportMaven"`
	SupportGradle         bool            `yaml:"supportGradle"`
	SupportDocker         bool            `yaml:"supportDocker"`
	SupportI18n           bool            `json:"supportI18n"`
	SupportSwagger        bool            `json:"supportSwagger"`
	SupportDataSource     string          `json:"supportDataSource"`
	HttpPort              string          `json:"httpPort"`
	Database              config.Database `json:"database"`
	Redis                 config.Redis    `json:"redis"`
	SupportRedis          bool            `json:"supportRedis"`
	JdbcDriverClass       string          `json:"jdbcDriverClass"`
	DBTypeMariadb         string          `json:"dbTypeMariadb"`
	DBTypeMysql           string          `json:"dbTypeMysql"`
	DBTypePostgresql      string          `json:"dbTypePostgresql"`
	DataSourceDruid       string          `json:"dataSourceDruid"`
	DataSourceHikariCP    string          `json:"dataSourceHikariCP"`
	SupportConfigTypeYaml bool            `json:"supportConfigTypeYaml"`
}

// 初始化项目数据
func (projectInfoDto SpringBootProjectInfoDto) Init() SpringBootProjectInfoDto {
	projectInfoDto.NowDate = time.Now().Format(config.NowTimeFormat)
	projectInfoDto.Author = config.ServerConfig.AuthorName
	projectInfoDto.HttpPort = config.ServerConfig.DefaultHttpPort
	projectInfoDto.Database = config.ServerConfig.Database
	projectInfoDto.Redis = config.ServerConfig.Redis
	projectInfoDto.SupportRedis = config.ServerConfig.Springboot.SupportRedis
	projectInfoDto.JdbcDriverClass = config.JDBCDriverClassNameMapping[projectInfoDto.Database.Type]
	projectInfoDto.DBTypePostgresql = config.DBTypePostgresql
	projectInfoDto.DBTypeMariadb = config.DBTypeMariadb
	projectInfoDto.DBTypeMysql = config.DBTypeMysql
	projectInfoDto.SupportMaven = config.ServerConfig.Springboot.SupportMaven
	projectInfoDto.SupportGradle = config.ServerConfig.Springboot.SupportGradle
	projectInfoDto.SupportDocker = config.ServerConfig.Springboot.SupportDocker
	projectInfoDto.SupportI18n = config.ServerConfig.Springboot.SupportI18n
	projectInfoDto.SupportSwagger = config.ServerConfig.Springboot.SupportSwagger
	projectInfoDto.SupportDataSource = config.ServerConfig.Springboot.SupportDataSource
	projectInfoDto.DataSourceDruid = config.DataSourceDruid
	projectInfoDto.DataSourceHikariCP = config.DataSourceHikariCP
	projectInfoDto.SupportConfigTypeYaml = config.ServerConfig.Springboot.SupportConfigTypeYaml
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
	projectInfoDto.ArtifactId = strings.Replace(projectInfoDto.ArtifactId, "-", ".", -1)
	projectInfoDto.PackageName = projectInfoDto.GroupId +
		"." + projectInfoDto.ArtifactId

	// package目录
	projectInfoDto.JavaPath = path.Join(projectInfoDto.ProjectName, config.JavaPath)
	projectInfoDto.ResourcePath = path.Join(projectInfoDto.ProjectName, config.JavaResourcePath)
	projectInfoDto.JavaTestPath = path.Join(projectInfoDto.ProjectName, config.JavaTestPath)
	packageNameList := strings.Split(projectInfoDto.PackageName, ".")
	for _, value := range packageNameList {
		projectInfoDto.JavaPath = path.Join(projectInfoDto.JavaPath, value)
		projectInfoDto.JavaTestPath = path.Join(projectInfoDto.JavaTestPath, value)
	}
	return projectInfoDto
}
func (projectInfoDto SpringBootProjectInfoDto) InitProject() {
	data, _ := json.MarshalIndent(projectInfoDto, "", "    ")
	fmt.Printf("%s\n", data)
	// 检测项目文件夹是否存在
	exist, err := fileUtil.PathExists(projectInfoDto.ProjectName)
	if err != nil {
		panic(err)
	}
	if exist {
		fmt.Printf("Project %s has Exist!\n", projectInfoDto.ProjectName)
		os.Exit(2)
	} else {
		fileUtil.PathMkDir(projectInfoDto.ProjectName)
	}
	// 检测package目录是否存在
	fileUtil.CheckDirAndMkdir(projectInfoDto.JavaPath)
	// 检查单元测试目录是否存在
	fileUtil.CheckDirAndMkdir(projectInfoDto.JavaTestPath)
	// 检测resource目录是否存在
	fileUtil.CheckDirAndMkdir(projectInfoDto.ResourcePath)
	// 检测model目录是否存在
	JavaCodeModelPath := path.Join(projectInfoDto.JavaPath, config.JavaModelPath)
	fileUtil.CheckDirAndMkdir(JavaCodeModelPath)
	// 检测mapper目录是否存在
	JavaCodeMapperPath := path.Join(projectInfoDto.JavaPath, config.JavaMapperPath)
	fileUtil.CheckDirAndMkdir(JavaCodeMapperPath)
	// 检测idleaf目录是否存在
	JavaCodeIdLeafPath := path.Join(projectInfoDto.JavaPath, config.JavaIdLeafPath)
	fileUtil.CheckDirAndMkdir(JavaCodeIdLeafPath)
	// 检测Util目录是否存在
	JavaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
	fileUtil.CheckDirAndMkdir(JavaCodeUtilPath)
	// 检测Config目录是否存在
	JavaCodeConfigPath := path.Join(projectInfoDto.JavaPath, config.JavaConfigPath)
	fileUtil.CheckDirAndMkdir(JavaCodeConfigPath)
	// config redis
	if projectInfoDto.SupportRedis {
		fileUtil.CheckDirAndMkdir(path.Join(JavaCodeConfigPath, config.JavaTemplateRedis))
	}
	// config swagger
	if projectInfoDto.SupportSwagger {
		fileUtil.CheckDirAndMkdir(path.Join(JavaCodeConfigPath, config.JavaTemplateSwaggerConfig))
	}
	// config I18n
	if projectInfoDto.SupportI18n {
		fileUtil.CheckDirAndMkdir(path.Join(JavaCodeConfigPath, config.JavaTemplateI18nConfig))
	}
	// config druid
	if config.DataSourceDruid == projectInfoDto.SupportDataSource {
		fileUtil.CheckDirAndMkdir(path.Join(JavaCodeConfigPath, config.JavaTemplateDruidConfig))
	}
	// config exception
	fileUtil.CheckDirAndMkdir(path.Join(JavaCodeConfigPath, config.JavaTemplateExceptionConfig))
	//
	// 检测Common目录是否存在
	JavaCodeCommonPath := path.Join(projectInfoDto.JavaPath, config.JavaCommonPath)
	fileUtil.CheckDirAndMkdir(JavaCodeCommonPath)
	// 检测mybatis目录是否存在
	mybatisPath := path.Join(projectInfoDto.ResourcePath, config.MybatisPath)
	fileUtil.CheckDirAndMkdir(mybatisPath)
	// 检查i18n目录是否存在
	if config.ServerConfig.Springboot.SupportI18n {
		i18nPath := path.Join(projectInfoDto.ResourcePath, config.JavaTemplateI18nProperties)
		fileUtil.CheckDirAndMkdir(i18nPath)
	}
	// 解析模板
	initProjectData(projectInfoDto)
}

func initProjectData(projectInfoDto SpringBootProjectInfoDto) {
	fileMapDtoList := []FileMapDto{
		// .project
		{path.Join(config.JavaTemplateInitPath, config.DotProjectFileName),
			path.Join(projectInfoDto.ProjectName, config.DotProjectFileName)},
		// Springboot Application
		{path.Join(config.JavaTemplateInitCodePath, config.JavaApplicationFileName),
			path.Join(projectInfoDto.JavaPath, config.JavaApplicationFileName)},
		// .gitignore
		{path.Join(config.JavaTemplateInitPath, config.GitIgnoreFileName),
			path.Join(projectInfoDto.ProjectName, config.GitIgnoreFileName)},
	}
	// README.md
	fileMapDtoList = append(fileMapDtoList, FileMapDto{
		path.Join(config.JavaTemplateInitPath, config.READMEFileName),
		path.Join(projectInfoDto.ProjectName, config.READMEFileName)})
	// Maven
	if projectInfoDto.SupportMaven {
		fileMapDtoList = append(fileMapDtoList, FileMapDto{
			path.Join(config.JavaTemplateInitPath, config.PomXmlFileName),
			path.Join(projectInfoDto.ProjectName, config.PomXmlFileName)})
	}
	// gradle
	if projectInfoDto.SupportGradle {
		fileMapDtoList = append(fileMapDtoList, FileMapDto{
			path.Join(config.JavaTemplateInitPath, config.GradleBuildFileName),
			path.Join(projectInfoDto.ProjectName, config.GradleBuildFileName)})
		fileMapDtoList = append(fileMapDtoList, FileMapDto{
			path.Join(config.JavaTemplateInitPath, config.GradleSettingFileName),
			path.Join(projectInfoDto.ProjectName, config.GradleSettingFileName)})
	}

	// Java Util包
	javaTemplateUtilPath := path.Join(config.JavaTemplateInitCodePath, config.JavaUtilPath)
	javaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
	fileMapDtoList = appendTemplateList(javaTemplateUtilPath, javaCodeUtilPath, fileMapDtoList)
	// SubDir holder
	fileMapDtoList = appendSubDirTemplateList(javaTemplateUtilPath, javaCodeUtilPath, fileMapDtoList)
	// I18n util
	if projectInfoDto.SupportI18n {
		javaTemplateUtilPath := path.Join(config.JavaTemplateInitCodePath, config.JavaTemplateI18nUtil)
		javaCodeUtilPath := path.Join(projectInfoDto.JavaPath, config.JavaUtilPath)
		fileMapDtoList = appendTemplateList(javaTemplateUtilPath, javaCodeUtilPath, fileMapDtoList)
	}

	// Java Config包
	javaTemplateConfigPath := path.Join(config.JavaTemplateInitCodePath, config.JavaConfigPath)
	javaCodeConfigPath := path.Join(projectInfoDto.JavaPath, config.JavaConfigPath)
	fileMapDtoList = appendTemplateList(javaTemplateConfigPath, javaCodeConfigPath, fileMapDtoList)
	// idLeaf
	javaTemplateModelPath := path.Join(config.JavaTemplateInitCodePath, config.JavaModelPath)
	javaCodeIdModelPath := path.Join(projectInfoDto.JavaPath, config.JavaModelPath)
	fileMapDtoList = appendTemplateList(javaTemplateModelPath, javaCodeIdModelPath, fileMapDtoList)
	javaTemplateMapperPath := path.Join(config.JavaTemplateInitCodePath, config.JavaMapperPath)
	javaCodeIdMapperPath := path.Join(projectInfoDto.JavaPath, config.JavaMapperPath)
	fileMapDtoList = appendTemplateList(javaTemplateMapperPath, javaCodeIdMapperPath, fileMapDtoList)
	javaTemplateIdLeafPath := path.Join(config.JavaTemplateInitCodePath, config.JavaIdLeafPath)
	javaCodeIdLeafPath := path.Join(projectInfoDto.JavaPath, config.JavaIdLeafPath)
	fileMapDtoList = appendTemplateList(javaTemplateIdLeafPath, javaCodeIdLeafPath, fileMapDtoList)
	// config swagger
	if projectInfoDto.SupportSwagger {
		fileMapDtoList = appendTemplateList(
			path.Join(javaTemplateConfigPath, config.JavaTemplateSwaggerConfig),
			path.Join(javaCodeConfigPath, config.JavaTemplateSwaggerConfig),
			fileMapDtoList)
	}
	// config I18n
	if projectInfoDto.SupportI18n {
		fileMapDtoList = appendTemplateList(
			path.Join(javaTemplateConfigPath, config.JavaTemplateI18nConfig),
			path.Join(javaCodeConfigPath, config.JavaTemplateI18nConfig),
			fileMapDtoList)
		fileMapDtoList = appendTemplateList(
			path.Join(javaTemplateConfigPath, config.JavaTemplateExceptionI18nConfig),
			path.Join(javaCodeConfigPath, config.JavaTemplateExceptionConfig),
			fileMapDtoList)
	} else {
		fileMapDtoList = appendTemplateList(
			path.Join(javaTemplateConfigPath, config.JavaTemplateExceptionConfig),
			path.Join(javaCodeConfigPath, config.JavaTemplateExceptionConfig),
			fileMapDtoList)
	}
	// config druid
	if config.DataSourceDruid == projectInfoDto.SupportDataSource {
		fileMapDtoList = appendTemplateList(
			path.Join(javaTemplateConfigPath, config.JavaTemplateDruidConfig),
			path.Join(javaCodeConfigPath, config.JavaTemplateDruidConfig),
			fileMapDtoList)
	}
	// Java Common包
	if projectInfoDto.SupportI18n {
		javaTemplateCommonPath := path.Join(config.JavaTemplateInitCodePath, config.JavaTemplateI18nCommon)
		javaCodeCommonPath := path.Join(projectInfoDto.JavaPath, config.JavaCommonPath)
		fileMapDtoList = appendTemplateList(javaTemplateCommonPath, javaCodeCommonPath, fileMapDtoList)
	} else {
		javaTemplateCommonPath := path.Join(config.JavaTemplateInitCodePath, config.JavaCommonPath)
		javaCodeCommonPath := path.Join(projectInfoDto.JavaPath, config.JavaCommonPath)
		fileMapDtoList = appendTemplateList(javaTemplateCommonPath, javaCodeCommonPath, fileMapDtoList)
	}
	// redis
	if projectInfoDto.SupportRedis {
		fileMapDtoList = appendTemplateList(
			path.Join(javaTemplateConfigPath, config.JavaTemplateRedis),
			path.Join(javaCodeConfigPath, config.JavaTemplateRedis),
			fileMapDtoList)
		// redis单元测试文件
		fileMapDtoList = appendTemplateList(config.JavaTemplateInitTestPath,
			projectInfoDto.JavaTestPath, fileMapDtoList)
	}
	// Resource文件
	fileMapDtoList = appendTemplateList(config.JavaTemplateResourcePath,
		projectInfoDto.ResourcePath, fileMapDtoList)
	if projectInfoDto.SupportConfigTypeYaml {
		fileMapDtoList = appendTemplateList(path.Join(config.JavaTemplateResourcePath, "yaml"),
			projectInfoDto.ResourcePath, fileMapDtoList)
	} else {
		fileMapDtoList = appendTemplateList(path.Join(config.JavaTemplateResourcePath, "default"),
			projectInfoDto.ResourcePath, fileMapDtoList)
	}
	// Mybatis
	mybatisTemplatePath := path.Join(config.JavaTemplateInitPath, config.MybatisPath)
	mybatisPath := path.Join(projectInfoDto.ResourcePath, config.MybatisPath)
	fileMapDtoList = appendMybatisTemplateList(mybatisTemplatePath, mybatisPath, fileMapDtoList)
	// I18n
	if projectInfoDto.SupportI18n {
		mybatisTemplatePath := path.Join(config.JavaTemplateInitPath, config.JavaTemplateI18nProperties)
		mybatisPath := path.Join(projectInfoDto.ResourcePath, config.JavaTemplateI18nProperties)
		fileMapDtoList = appendTemplateList(mybatisTemplatePath, mybatisPath, fileMapDtoList)
	}
	dataList, err := json.MarshalIndent(fileMapDtoList, "", "    ")
	fmt.Printf("%s\n", dataList)
	if err != nil {
		panic(err)
	}
	// 解析模板
	for _, value := range fileMapDtoList {
		templateUtil.ParseTemplate(value.TplDstPath, value.TplSrcPath, projectInfoDto)
	}
}
