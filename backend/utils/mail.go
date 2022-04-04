package utils

import (
	"backend/config"
	"fmt"
	"net/smtp"
	"strings"
)

func SendToMail(user, sendUserName, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendEmail(toEmail string, subject string, content string) {
	//user := "windypath_system@126.com"
	//password := "GAICYEMOHNMZFTSI"
	//host := "smtp.126.com:25"
	user := config.EmailConfig.User
	password := config.EmailConfig.Password
	host := config.EmailConfig.Host
	//to := "289348588@qq.com;3483131756@qq.com"
	to := toEmail

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>` + content +
		`
			
		</body>
		</html>`

	sendUserName := "赛事预约提醒系统" //发送邮件的人名称
	fmt.Println("send email")
	err := SendToMail(user, sendUserName, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("SendEmail mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("SendEmail mail success!")
	}

}
