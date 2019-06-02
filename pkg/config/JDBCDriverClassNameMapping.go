package config

var DBTypeMariadb = "mariadb"
var DBTypeMysql = "mysql"
var DBTypePostgresql = "postgresql"

var DataSourceDruid = "druid"
var DataSourceHikariCP = "HikariCP"
var JDBCDriverClassNameMapping = map[string]string{
	DBTypeMariadb:    "org.mariadb.jdbc.Driver",
	DBTypeMysql:      "com.mysql.jdbc.Driver",
	DBTypePostgresql: "org.postgresql.Driver",
}
