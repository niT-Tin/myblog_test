package global

import (
	"myblog/initdb"
	"myblog/initdb/mysql"
)

var (
	e             initdb.IError
	mysqlHost     string
	mysqlPort     string
	mysqlUser     string
	mysqlPassword string

	mongoDBName         string
	mongoCollectionName string

	redisHost string
	redisPort string
)

func init() {
	ErrorHandler := initdb.NewError(mysql.DB)
	e = ErrorHandler
}

// GetErrorHandler 获取错误处理全局对象
func GetErrorHandler() initdb.IError {
	return e
}
