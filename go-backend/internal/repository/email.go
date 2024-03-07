package repository

import (
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

func SendEmail(emailAddress, username, question, answer, comment string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "SERVE <support@agileserve.org.cn>")
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", "CALLMe: 助教/老师回答了你的问题")
	m.SetBody("text/html", "亲爱的 <b>"+username+"</b>，<br><br>"+
		"关于你的提问: <b>"+question+"</b><br><br>"+
		"及CALLMe给出的答案: <b>"+answer+"</b><br><br>"+
		"助教/老师给出了相关评论: <b>"+comment+"</b><br><br><br>"+
		"感谢你使用CALLMe服务！我们期待继续为你提供帮助。<br><br>"+
		"CALLMe团队")

	d := gomail.NewDialer("smtp.exmail.qq.com", 587, "support@agileserve.org.cn", "HuMeaXb2DDDCQGA8")

	maxAttempts := 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		// 蓝色字体
		log.Println("\033[34mSending email\033[0m to", emailAddress, "on attempt", attempt)
		if err := d.DialAndSend(m); err != nil {
			// 红色字体
			log.Println("\033[31mFailed to send email\033[0m to", emailAddress, "on attempt", attempt, "with error:", err)
			if attempt < maxAttempts {
				log.Println("Retrying in", attempt, "minute(s)...")
				time.Sleep(time.Duration(attempt) * time.Minute)
				continue
			}
		} else {
			// 绿色字体
			log.Println("\033[32mEmail sent\033[0m to", emailAddress, "on attempt", attempt)
			break
		}
	}
}
