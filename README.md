# aliyun-go-smtpd
aliyun-go-smtpd是go语言实现在阿里云服务器实现email的收发服务的应用,使用smtp协议收取邮件,并将邮件内容以日志的形式存储起来,有开发能力的开发者也可以对收取的邮件进行本地数据库存储和对邮件内容抓取(比如:验证码,网址,手机号等)特殊字符串的中间件预留了处理接口.

### 搭建要求

* 拥有域名解析权限的域名一个(域名并且已经备案)
* 阿里云服务器一台

### 搭建环境:

	1.go version:1.16+
	2.`receives/`需要有写入权限。
	3.启动方式：go run main.go
	4.编译：go build main.go
	5.后台运行：nohup go run main.go &>/dev/null &
	6.编译后后台运行：nohup ./main &>/dev/null &
	7.特别说明:服务器要开启端口号`25`的访问权限,阿里云服务器默认是关闭的,可以联系客服开启,并且服务器防火墙也不能拦截`25端口`.程序运行之后可以执行`netstat -lntp 0.0.0.0:25`查看是否有25端口的监听.
	
	
### 域名解析说明

* 域名添加值为`mail`的A记录指向你的阿里云服务器Ip.
* 再添加值为`mail.demo.com(你的域名)`的MX记录.

如下图:

![demo-dns](https://github.com/langzishen/aliyun-go-smtpd/blob/main/demo-dns.png)