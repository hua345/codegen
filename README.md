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
### 接口生成
主要根据URL设计和方法来生成对应的接口
```bash
./codegen.exe -a hello.fang -m addArticle -url addArticle
./codegen.exe -a hello.fang -m addUser -url user
./codegen.exe -a hello.fang -m getUser -url user -httpMethod get
./codegen.exe -a hello.fang -m getDetail -url user/{id} -httpMethod get
./codegen.exe -a hello.fang -m deleteUser -url user -httpMethod delete
```