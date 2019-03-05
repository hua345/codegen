// generateCode project generateCode.go
package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Code Generate version:1.0.0
Usage: ./generateCode.exe -init`)
	flag.PrintDefaults()
}
func main() {
	//jsonCode.ReadJsonData()
flag.Usage = usage
flag.Parse()
}
