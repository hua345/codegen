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
	"codegen/models"
	"codegen/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)
var artifactId string
// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [-a ArtifactId]",
	Short: "Springboot初始化工程",
	Long: `Springboot初始化工程. For example:
codegen init -a ArtifactId`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Setup(cfgFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Inside subCmd Run with args: ", args)
		if len(artifactId) == 0 {
			artifactId = config.ServerConfig.Springboot.ArtifactId
		}
		projectInfoDto := models.SpringBootProjectInfoDto{GroupId: config.ServerConfig.Springboot.GroupId, ArtifactId: artifactId}
		projectInfoDto = projectInfoDto.Init()
		projectInfoDto.InitProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&artifactId, "artifactId", "a", "", "ArtifactID 格式：产品线名-模块名。")
}
