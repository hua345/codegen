> 一千个人心中有一千个哈姆雷特，一千个架构师心中有一千种完美架构。
>
> 通过http/rpc协议，将自己喜欢的几种语言协作起来，发挥各自优点。

![badge.svg?branch=master](https://github.com/hua345/codegen/workflows/Go/badge.svg?branch=master)

### 编译安装

```bash
# 下载源码
git clone https://github.com/hua345/codegen.git
# 安安装go-bindata
go get github.com/spf13/cobra
# github.com/go-yaml/yaml -> gopkg.in/yaml.v2
go get github.com/go-yaml/yaml
go get github.com/sirupsen/logrus
go get github.com/go-bindata/go-bindata/...
# 打包静态资源
go generate
# 编译程序
go build
# 安装程序
go install
```

### 1. 查看帮助

```bash
 codegen.exe -h
 Code Generate Springboot/SpringCloud/Gin/Express Helper

Usage:
  codegen [command]

Available Commands:
  api         Springboot接口生成
  cloud       SpringCloud代码生成
  config      配置文件命令
  express     Node express代码生成
  gin         Golang Gin代码生成
  help        Help about any command
  init        Springboot初始化工程

Flags:
  -f, --config string   读取配置文件(默认./codegen.yaml)
  -h, --help            help for codegen

Use "codegen [command] --help" for more information about a command.
```

### 2. 初始化配置文件

```bash
codegen.exe config init
```

会在当前目录下生成`codegen.yaml`配置文件

#### 3.1 阿里GroupId和ArtifactId规范

- GroupID 格式：com.{公司/BU }.业务线.[子业务线]，最多 4 级。

> 说明：{公司/BU} 例如：alibaba/taobao/tmall/aliexpress 等 BU 一级；子业务线可选。
正例：com.taobao.jstorm 或 com.alibaba.dubbo.register

- ArtifactID 格式：产品线名-模块名。语义不重复不遗漏，先到中央仓库去查证一下。
  
> 正例：dubbo-client / fastjson-api / jstorm-tool

#### 3.2 修改配置文件`codegen.yaml`

```yaml
defaultHttpMethod: post
defaultHttpPort: 8080
apiBaseUrl: api/v1
authorName: chenjianhua
database:
  # postgresql/mariadb/mysql/
  type: mariadb
  host: 192.168.137.128:3306
  databaseName: db_example
  username: springuser
  password: 123456
redis:
  host: 192.168.137.128
  port: 6379
  # Redis默认情况下有16个分片，这里配置具体使用的分片，默认是0
  database: 0
  password:
  # 连接超时时间（毫秒）
  timeout: 1000
  # 连接池中的最大空闲连接 默认 8
  maxIdle: 8
  # 连接池最大连接数（使用负值表示没有限制） 默认 8
  maxActive: 8


springboot:
  groupId: com.github
  artifactId: code-admin
  supportRedis: true
  #shiro, security, spring-session, redis-session, none
  userService: none
  supportMaven: true
  # 默认配置文件类型是properties,也支持yaml
  supportConfigTypeYaml: false
  supportGradle: true
  supportDocker: true
  # 国际化
  supportI18n: true
  # 数据源: druid/HikariCP
  supportDataSource: druid
  supportSwagger: true
```

#### 3.3 工程初始化

```bash
codegen.exe init
```

### 4.1 查看帮助

```bash
$ codegen.exe api -h
Springboot接口生成工具

Usage:
  codegen api [flags]

Examples:
./codegen.exe api -m methodName -u url [-a ArtifactId]
[-r requestMethod] [-g GroupId] [-d Description]

Flags:
  -a, --ArtifactId string      ArtifactID 格式：产品线名-模块名。
  -d, --description string     接口描述
  -g, --groupId string         GroupID 格式：com.{公司/BU }.业务线.[子业务线]
  -h, --help                   help for api
  -m, --methodName string      Controller类中的方法名称
  -r, --requestMethod string   http请求方式Get/Post/Delete，默认: Get
  -u, --urlPath string         Url路径

Global Flags:
  -f, --config string   读取配置文件(默认./codegen.yaml)
```

#### 4.2 接口生成

主要根据URL设计和方法来生成对应的接口

```bash
codegen.exe api -m addArticle -u addArticle
codegen.exe api -m getUser -u user
codegen.exe api -m addUser -u user -r post
codegen.exe api -m getUserDetail -u user/{id} -r get
codegen.exe api -m deleteUser -u user -r delete
```
