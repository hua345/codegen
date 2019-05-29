### 编译
```bash
go get github.com/hua345/codegen
```
### 查看帮助
```bash
 ./codegen.exe -h
```
### 初始化工程
```bash
./codegen.exe -init  -a ArtifactId [-group GroupId]
```
#### 阿里GroupId和ArtifactId规范
- GroupID 格式：com.{公司/BU }.业务线.[子业务线]，最多 4 级。
> 说明：{公司/BU} 例如：alibaba/taobao/tmall/aliexpress 等 BU 一级；子业务线可选。
正例：com.taobao.jstorm 或 com.alibaba.dubbo.register
- ArtifactID 格式：产品线名-模块名。语义不重复不遗漏，先到中央仓库去查证一下。
> 正例：dubbo-client / fastjson-api / jstorm-tool


### 接口生成
主要根据URL设计和方法来生成对应的接口
```bash
./codegen.exe -a hello-fang -m addArticle -url addArticle
./codegen.exe -a hello-fang -m addUser -url user
./codegen.exe -a hello-fang -m getUser -url user -httpMethod get
./codegen.exe -a hello-fang -m getDetail -url user/{id} -httpMethod get
./codegen.exe -a hello-fang -m deleteUser -url user -httpMethod delete
```