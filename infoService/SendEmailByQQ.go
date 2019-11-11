package main

import (
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func SendEmail(mailTo []string, subject string, content string) error {
	mailConn := map[string]string{
		"user": "1759540235@qq.com",
		"pass": "rlaueuhkmpradjga",
		"host": "smtp.qq.com",
		"port": "465",
	}
	port, err := strconv.Atoi(mailConn["port"])
	if err != nil {
		log.Fatal(err)
		return err
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">")
	mail.SetHeader("To", mailTo...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", "content")
	dialer := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err = dialer.DialAndSend(mail)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("邮件发送成功")
	return err
}

func main() {
	//定义收件人
	mailTo := []string{
		"13263397018@163.com",
	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "Good"
	SendEmail(mailTo, subject, body)
}
