package config

var DBTypeMariadb = "mariadb"
var DBTypeMysql = "mysql"
var DBTypePostgresql = "postgresql"

var JDBCDriverClassNameMapping = map[string]string{
	DBTypeMariadb:    "com.mariadb.jdbc.Driver",
	DBTypeMysql:      "com.mysql.jdbc.Driver",
	DBTypePostgresql: "org.postgresql.Driver",
}
