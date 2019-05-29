package config

import (
	"codegen/pkg/util"
	"fmt"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type Server struct {
	DefaultHttpMethod string            `yaml:"defaultHttpMethod"`
	DefaultHttpPort   string            `yaml:"defaultHttpPort"`
	ApiBaseUrl        string            `yaml:"apiBaseUrl"`
	AuthorName        string            `yaml:"authorName"`
	Database          Database          `yaml:"database"`
	Redis             Redis             `yaml:"redis"`
	Springboot        SpringbootSetting `yaml:"springboot"`
}

func newServer() Server {
	return Server{
		DefaultHttpMethod: "com.github",
		DefaultHttpPort:   "api/v1",
		ApiBaseUrl:        "api/v1",
		AuthorName:        "learn",
	}
}

type SpringbootSetting struct {
	GroupId    string `yaml:"groupId"`
	ArtifactId string `yaml:"artifactId"`
}

func newSpringbootSetting() SpringbootSetting {
	return SpringbootSetting{
		GroupId:    "com.github",
		ArtifactId: "learn",
	}
}

type Database struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Redis struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

type GinSetting struct {
}

var HttpMethodMapping = map[string]string{
	"GET":    "GetMapping",
	"POST":   "PostMapping",
	"PUT":    "PutMapping",
	"DELETE": "DeleteMapping",
}

var ServerConfig Server
var DefaultSettingFile = DefaultConfigFile

// Setup initialize the configuration instance
func Setup(configPath string) {
	ServerConfig = newServer()
	path := DefaultSettingFile
	if len(configPath) > 0 {
		path = configPath
	}
	configExist, _ := util.PathExists(path)
	if true != configExist {
		log.Println("配置文件" + DefaultConfigFile + "没有找到!")
		log.Println("初始化配置文件命令: codegen config init")
		os.Exit(-1)
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// set default
	ServerConfig.Springboot = newSpringbootSetting()
	err = yaml.Unmarshal(yamlFile, &ServerConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
}
