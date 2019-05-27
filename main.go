// generateCode project main.go
package main

import (
	"flag"
	"fmt"
	"codegen/models"
	"os"
	"codegen/pkg/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

}
func usage() {
	fmt.Fprintf(os.Stderr, `Code Generate version:1.0.0
Usage: ./generateCode.exe -init  -a ArtifactId [-group GroupId] #初始化工程
       ./generateCode.exe -a ArtifactId -m methodName -url url
       [-httpMethod GET/POST] [-group GroupId] [-baseUrl baseUrl] #生成单个接口
`)
	flag.PrintDefaults()
}
func main() {
	servicePtr := flag.String("a", "", "项目名称，比如:demo、demo.code")
	methodPtr := flag.String("m", "", "方法名称")
	groupNamePtr := flag.String("group", config.ServerConfig.Springboot.GroupId, "组织名称，比如:com.github")
	restfulUrlPtr := flag.String("url", "", "URL路径，比如:user")
	baseUrlPtr := flag.String("baseUrl", config.ServerConfig.Springboot.ApiBaseUrl, "URL路径，比如:user")
	httpMethodPtr := flag.String("httpMethod", "", "Http Method，比如:get, post,put,delete")
	initPtr := flag.Bool("init", false, "初始化工程")
	flag.Usage = usage
	flag.Parse()

	config.Setup("")
	if *initPtr {
		if len(*servicePtr) == 0 {
			flag.Usage()
			os.Exit(0)
		}
		projectInfoDto := models.ProjectInfoDto{GroupId: *groupNamePtr, ArtifactId: *servicePtr}
		projectInfoDto = projectInfoDto.Init()
		projectInfoDto.InitProject()
	} else {
		if len(*servicePtr) == 0 || len(*methodPtr) == 0 || len(*restfulUrlPtr) == 0 {
			flag.Usage()
			os.Exit(0)
		}
		projectInfoDto := models.ProjectInfoDto{GroupId: *groupNamePtr, ArtifactId: *servicePtr}
		projectInfoDto = projectInfoDto.Init()
		restfulApiDto := models.RestfulApiDto{
			HttpMethod:  *httpMethodPtr,
			MethodName:  *methodPtr,
			BaseUrl:     *baseUrlPtr,
			RestfulUrl:  *restfulUrlPtr,
			ProjectInfo: projectInfoDto}
		restfulApiDto = restfulApiDto.Init()
		restfulApiDto.GenerateCode()
	}
}
