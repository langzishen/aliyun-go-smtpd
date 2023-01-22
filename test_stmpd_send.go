package main

import (
	"fmt"
	"net/smtp"
)

func stmpd_send() {
	c, e := smtp.Dial("127.0.0.1:25")
	defer c.Close()
	//fmt.Println(c, ee)
	if e != nil {
		fmt.Println(e)
	}

	c.Mail("from")
	c.Rcpt("to")
	wc, _ := c.Data()
	fmt.Fprintf(wc, "data")
	wc.Close()
	c.Quit()
}
