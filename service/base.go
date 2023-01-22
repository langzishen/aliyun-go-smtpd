package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
)

var (
	db_user          = "root"
	db_user_password = "chenhc"
	db_host          = "localhost"
	db_port          = 3306
	db_name          = "go-smtpd"
)

//dsn := "用户名:密码@tcp(127.0.0.1:端口号)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
var Dsn = db_user + ":" + db_user_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&parseTime=True"

func InitDB() *gorm.DB {
	Db, Err := gorm.Open(mysql.Open(Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	logFile, _ := os.OpenFile("./log_err/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logger := log.New(logFile, "err_Db:", log.Ldate|log.Ltime|log.Lshortfile)
	if Err != nil {
		//panic(Err.Error())
		//fmt.Println(Err.Error())
		logger.Println(Err.Error())
	}
	return Db
}

type Base struct {
	ModelName string
}
