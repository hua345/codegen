### 编译
```bash
go get github.com/hua345/codegen
```
### 1. 查看帮助
```bash
 ./codegen.exe -h
```
### 2. 初始化配置文件
```bash
./codegen.exe config init
```
#### 3.1 阿里GroupId和ArtifactId规范
- GroupID 格式：com.{公司/BU }.业务线.[子业务线]，最多 4 级。
> 说明：{公司/BU} 例如：alibaba/taobao/tmall/aliexpress 等 BU 一级；子业务线可选。
正例：com.taobao.jstorm 或 com.alibaba.dubbo.register
- ArtifactID 格式：产品线名-模块名。语义不重复不遗漏，先到中央仓库去查证一下。
> 正例：dubbo-client / fastjson-api / jstorm-tool

#### 3.2 修改配置文件`codegen.yaml`
```
defaultHttpMethod: get
defaultHttpPort: 8080
apiBaseUrl: api/v1
authorName: chenjianhua
springboot:
  groupId: com.github
  artifactId: hello-golang
```
#### 3.3 工程初始化
```
$ ./codegen.exe init -h
Springboot初始化工程. For example:
codegen init -a ArtifactId

Usage:
  codegen init [-a ArtifactId] [flags]

Flags:
  -a, --artifactId string   ArtifactID 格式：产品线名-模块名。
  -h, --help                help for init

Global Flags:
  -f, --config string   读取配置文件(默认./codegen.yaml)
```
```
./codegen.exe init
```
### 4. 接口生成
```
$ ./codegen.exe api -h
Springboot接口生成工具

Usage:
  codegen api [flags]

Examples:
./codegen.exe -m methodName -u url [-a ArtifactId]
[-r requestMethod] [-g GroupId] [-d Description] [-baseUrl baseUrl]

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
主要根据URL设计和方法来生成对应的接口
```bash
./codegen.exe api -m addArticle -u addArticle
./codegen.exe api -m getUser -u user
./codegen.exe api -m addUser -u user -r post
./codegen.exe api -m getUserDetail -u user/{id} -r get
./codegen.exe api -m deleteUser -u user -r delete
```
