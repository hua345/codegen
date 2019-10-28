// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/hua345/codegen/models"
	"github.com/hua345/codegen/pkg/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var (
	springbootHttpMethod  string
	springbootRestfulUrl  string
	springbootDescription string
	springbootGroupName   string
	springbootMethodName  string
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Springboot接口生成",
	Long:  "Springboot接口生成工具",
	Example: `./codegen.exe -m methodName -u url [-a ArtifactId]
[-r requestMethod] [-g GroupId] [-d Description] [-baseUrl baseUrl]`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Setup(cfgFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		checkFlag(cmd)
		springbootHttpMethod = strings.ToUpper(springbootHttpMethod)
		if !(config.HttpMethodGet == springbootHttpMethod ||
			config.HttpMethodPost == springbootHttpMethod ||
			config.HttpMethodUpdate == springbootHttpMethod ||
			config.HttpMethodPut == springbootHttpMethod ||
			config.HttpMethodDelete == springbootHttpMethod) {
			log.Println("支持的http请求方法:" + config.HttpMethodGet + "\\" +
				config.HttpMethodPost + "\\" + config.HttpMethodPut + "\\" + config.HttpMethodDelete)
			os.Exit(-1)
		}
		projectInfoDto := models.SpringBootProjectInfoDto{GroupId: springbootGroupName, ArtifactId: artifactId}
		projectInfoDto = projectInfoDto.Init()
		restfulApiDto := models.SpringBootRestfulApiDto{
			HttpMethod:  springbootHttpMethod,
			MethodName:  springbootMethodName,
			BaseUrl:     config.ServerConfig.ApiBaseUrl,
			RestfulUrl:  springbootRestfulUrl,
			ProjectInfo: projectInfoDto,
			Description: springbootDescription}
		restfulApiDto = restfulApiDto.Init()
		restfulApiDto.GenerateCode()
	},
}

func checkFlag(cmd *cobra.Command) {
	if len(springbootMethodName) == 0 {
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
		log.Println("需要参数-m methodName，Controller类中的方法名称")
		os.Exit(-1)
	}
	if len(springbootRestfulUrl) == 0 {
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
		log.Println("需要参数-u url, Url路径")
		os.Exit(-1)
	}
	if len(artifactId) == 0 {
		artifactId = config.ServerConfig.Springboot.ArtifactId
	}
	if len(springbootHttpMethod) == 0 {
		springbootHttpMethod = config.ServerConfig.DefaultHttpMethod
	}
	if len(springbootGroupName) == 0 {
		springbootGroupName = config.ServerConfig.Springboot.GroupId
	}
}
func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	apiCmd.Flags().StringVarP(&artifactId, "ArtifactId", "a", "", "ArtifactID 格式：产品线名-模块名。")
	apiCmd.Flags().StringVarP(&springbootMethodName, "methodName", "m", "", "Controller类中的方法名称")
	apiCmd.Flags().StringVarP(&springbootHttpMethod, "requestMethod", "r", "", "http请求方式Get/Post/Delete，默认: Get")
	apiCmd.Flags().StringVarP(&springbootRestfulUrl, "urlPath", "u", "", "Url路径")
	apiCmd.Flags().StringVarP(&springbootDescription, "description", "d", "", "接口描述")
	apiCmd.Flags().StringVarP(&springbootGroupName, "groupId", "g", "", "GroupID 格式：com.{公司/BU }.业务线.[子业务线]")
}
