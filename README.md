### 查看帮助
```bash
 ./generateCode.exe -h
```
### 初始化工程
```bash
./generateCode.exe -init  -a ArtifactId [-group GroupId]
```
### 接口生成
主要根据URL设计和方法来生成对应的接口
```bash
./generateCode.exe -a hello.fang -m addArticle -url addArticle
./generateCode.exe -a hello.fang -m addUser -url user
./generateCode.exe -a hello.fang -m getUser -url user -httpMethod get
./generateCode.exe -a hello.fang -m getDetail -url user/{id} -httpMethod get
./generateCode.exe -a hello.fang -m deleteUser -url user -httpMethod delete
```