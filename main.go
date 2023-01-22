package main

import (
	"encoding/json"
	"go-smtpd/service"
	"go-smtpd/smtpd"
	"log"
	"net"
	"os"
	"time"
)

var DbType = "" //数据库存储类型,如果为空不进行数据库存储, 可能值:mysql,redis

//处理接受的邮件 中间件
func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	logMailReceive(from, to, data)

	if DbType != "" {
		if DbType == "mysql" { //如果需要存储到数据库可以参考receives中的写法去写数据存储到数据库
			new(service.Receives).Add(from, to, string(data), "")
		}
	}
	return nil
}

//将接受的邮件信息写入日志
func logMailReceive(from string, to []string, data []byte) {
	fileName := time.Now().Format("20060102150405") + ".log"
	logFile, _ := os.OpenFile("./receives/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	map_data := map[string]string{"from": from, "to": to[0], "content": string(data)}
	datas, _ := json.Marshal(map_data)
	logFile.WriteString(string(datas))
}

func rcptHandler(remoteAddr net.Addr, from string, to string) bool {
	//domain = getDomain(to)
	//return domain == "mail.example.com"
	return true
}

//监听服务
func ListenAndServe(addr string, handler smtpd.Handler, rcpt smtpd.HandlerRcpt) error {
	srv := &smtpd.Server{
		Addr:        addr,
		Handler:     handler,
		HandlerRcpt: rcpt,
		Appname:     "MyServerApp",
		Hostname:    "",
	}
	return srv.ListenAndServe()
}

func main() {
	//err := ListenAndServe("127.0.0.1:25", mailHandler, rcptHandler)     //监听本机的25端口
	err := ListenAndServe("0.0.0.0:25", mailHandler, rcptHandler) //监听所有来源的25端口
	logFile, _ := os.OpenFile("./log_err/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logger := log.New(logFile, "err_ListenAndServe:", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		logger.Println(err.Error())
	}
}
