package main

import (
	"net/smtp"
	"strings"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	auth := smtp.PlainAuth("", "1041855188@qq.com", "密码", "smtp.qq.com")
	to := []string{"1339796113@qq.com"}
	nickname := "test"
	user := "1041855188@qq.com"
	subject := "test mail"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
	fmt.Println(time.Now())
}
