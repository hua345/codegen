package config

var DBTypeMariadb = "mariadb"
var DBTypeMysql = "mysql"
var DBTypePostgresql = "postgresql"

var JDBCDriverClassNameMapping = map[string]string{
	DBTypeMariadb:    "org.mariadb.jdbc.Driver",
	DBTypeMysql:      "com.mysql.jdbc.Driver",
	DBTypePostgresql: "org.postgresql.Driver",
}
