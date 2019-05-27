package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"codegen/pkg/util"
	"io/ioutil"
	"time"
	"gopkg.in/yaml.v2"
)

type Server struct {
	DefaultHttpMethod string            `yaml:"defaultHttpMethod"`
	DefaultHttpPort   string            `yaml:"defaultHttpPort"`
	AuthorName        string            `yaml:"authorName"`
	LangType          string            `yaml:"langType"`
	Database          Database          `yaml:"database"`
	Redis             Redis             `yaml:"redis"`
	Springboot        SpringbootSetting `yaml:"springboot"`
}

func newServer() Server {
	return Server{
		DefaultHttpMethod: "com.github",
		DefaultHttpPort:   "api/v1",
		AuthorName:        "learn",
		LangType:          "java",
	}
}

type SpringbootSetting struct {
	GroupId    string `yaml:"groupId"`
	ApiBaseUrl string `yaml:"apiBaseUrl"`
	ArtifactId string `yaml:"artifactId"`
}

func newSpringbootSetting() SpringbootSetting {
	return SpringbootSetting{
		GroupId:    "com.github",
		ApiBaseUrl: "api/v1",
		ArtifactId: "learn",
	}
}

type Database struct {
	Type     string `yaml:"type"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DbName   string `yaml:"dbName"`
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
var DefaultSettingFile = "codegen.yaml"
// Setup initialize the configuration instance
func Setup(configPath string) {
	ServerConfig = newServer()
	path := DefaultSettingFile
	if len(configPath) > 0 {
		path = configPath
	}
	configExist, _ := util.PathExists(path)
	if true != configExist {
		log.WithFields(log.Fields{
			"Path": path,
		}).Error("Config File Not Found!")
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
	log.Info("ServerConfig:", ServerConfig)
}
