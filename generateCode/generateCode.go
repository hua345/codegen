// generateCode project generateCode.go
package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Code Generate version:1.0.0
Usage: ./generateCode.exe -init  -a ArtifactId [-group GroupId] #初始化工程
       ./generateCode.exe -a ArtifactId -m methodName 
       [-group GroupId] #生成单个接口
`)
	flag.PrintDefaults()
}
func main() {
	//jsonCode.ReadJsonData()
	servicePtr := flag.String("a", "", "项目名称，比如:demo、demo.code")
	methodPtr := flag.String("m", "", "方法名称")
	groupNamePtr := flag.String("group", DefaultGroupName, "组织名称，比如:com.github")
	initPtr := flag.Bool("init", false, "初始化工程")
	flag.Usage = usage
	flag.Parse()

	if *initPtr {
		if len(*servicePtr) == 0 {
			flag.Usage()
			os.Exit(0)
		}
		fmt.Println("GroupId: ", *groupNamePtr)
		fmt.Println("init project: ", *servicePtr)

	} else {
		if len(*servicePtr) == 0 || len(*methodPtr) == 0 {
			flag.Usage()
			os.Exit(0)
		}
		fmt.Println("GroupId: ", *groupNamePtr)
		fmt.Println(*servicePtr + " project generate Api: " , *methodPtr)

	}
}
