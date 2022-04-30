package check

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {

	mailConn := map[string]string{
		"user": "chaiyd.cn@gmail.com",
		"pass": "xxxxxxxx",
		"host": "smtp.gmail.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "devops"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
func main() {
	//定义收件人
	mailTo := []string{
		"chaiyd.cn@gmail.com",
	}
	//邮件主题为"Hello"
	subject := "Hello by golang gomail from smtp.gmail.com"
	// 邮件正文
	body := "Hello,by gomail sent"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")

}
