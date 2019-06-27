package config

var DefaultConfigContent = `defaultHttpMethod: post
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
springboot:
  groupId: com.github
  artifactId: code-admin
  supportRedis: true
  supportMaven: true
  supportGradle: true
  supportDocker: true
  # 国际化
  supportI18n: true
  # 数据源: druid/HikariCP
  supportDataSource: druid
  supportSwagger: true`
