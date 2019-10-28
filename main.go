//go:generate go-bindata -o=asset/asset.go -pkg=asset public/...
package main

import (
	"github.com/hua345/codegen/cmd"
	log "github.com/sirupsen/logrus"
	"os"
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
func main() {
	cmd.Execute()
}
