#### {{.ProjectName}}
Springboot项目

### 构建
{{ if .SupportMaven }}
#### Maven构建
```
mvn clean install
```
{{ end }}
{{ if .SupportGradle }}
#### Gradle构建
```
gradle build
```
{{ end }}
