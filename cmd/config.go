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
	"github.com/hua345/codegen/pkg/config"
	"github.com/hua345/codegen/pkg/fileUtil"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置文件命令",
	Long:  "修改codegen配置文件",
	Example: `codegen config init
codegen config list`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
	},
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置文件",
	Long: `初始化配置文件. For example:
codegen config init`,
	Run: func(cmd *cobra.Command, args []string) {
		configExist, _ := fileUtil.PathExists(config.DefaultConfigFile)
		if true == configExist {
			log.WithFields(log.Fields{
				"Config File": config.DefaultConfigFile,
			}).Error("Config File is Exist!")
			os.Exit(-1)
		}
		fileUtil.WriteFileWithIoUtil(config.DefaultConfigFile, config.DefaultConfigContent)
		log.Info("Init config file success")
	},
}
var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出配置文件内容",
	Long: `列出配置文件内容. For example:
codegen config list`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Setup(cfgFile)
		data, _ := json.MarshalIndent(config.ServerConfig, "", "    ")
		fmt.Printf("%s\n", data)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configListCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
