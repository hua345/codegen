### 初始化工程
```bash
./generateCode.exe -init  -a ArtifactId [-group GroupId]
```
### 接口生成
主要根据URL设计和方法来生成对应的接口
```bash
project baseUrl = /api/v1/hello
GET /users/{name} method getUser
=> Controller Users method getUser
POST /users method addUser
=> Controller Users method addUser
POST /addUser [method addUser]
=> Controller AddUser method addUser
```
```bash
./generateCode.exe -a ArtifactId -m methodName -url url
[-httpMethod GET/POST] [-group GroupId] [-baseUrl baseUrl]
```