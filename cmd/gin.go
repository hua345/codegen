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
	"fmt"
	"github.com/spf13/cobra"
)

// ginCmd represents the gin command
var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "Golang Gin代码生成",
	Long: `Golang Gin代码生成. For example:
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Setup(cfgFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Gin代码生成: TODO")
	},
}

func init() {
	// rootCmd.AddCommand(ginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
