### 初始化工程
```bash
./generateCode.exe -init  -a ArtifactId [-group GroupId]
```
### 接口生成
主要根据URL设计和方法来生成对应的接口
```bash
project baseUrl = /api/v1/hello
GET article/details/{id} method getDetails
=> Controller Article method getDetails
POST /article method addArticle
=> Controller Article method addArticle
POST /addArticle
=> Controller AddArticle method addArticle
```
```bash
./generateCode.exe -a ArtifactId -m methodName -url url
[-httpMethod GET/POST] [-group GroupId] [-baseUrl baseUrl]
```