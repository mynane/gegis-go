package utils

import (
	"log"
	"github.com/robfig/cron"
	"github.com/go-gomail/gomail"
	"strings"
)

type Message struct {
	email		string "to"
	subject 	string "subject"
	contentType string "text/html"
	body 		string "body"
}

const (
	HOST 		= "smtp.qq.com" 		// 发送邮件服务器
	PORT 		= 456 					// 端口
	USERNAME 	= "" 	// 发件人账号
	PASSWORD	= "" 	// 发件人密码
)

func Send(message *Message) error {
	sendTo := strings.Split(message.email, ";")
	done := make(chan error, 1024)
	d := gomail.NewDialer(HOST, PORT, USERNAME, PASSWORD)
	go func() {
		defer close(done)
		for _, v := range sendTo {
			m := gomail.NewMessage()
			m.SetHeaders(map[string][]string{
				"From":    {m.FormatAddress(USERNAME, "极客圈")},
				"To":      {v},
				"Subject": {message.subject},
			})
			m.SetAddressHeader("From", v /*"发件人地址"*/, "发件人") // 发件人
			m.SetHeader("To", m.FormatAddress(message.email, "收件人")) // 收件人
			//m.SetHeader("Cc",
			//	m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送
			//m.SetHeader("Bcc", m.FormatAddress("xxxx@gmail.com", "收件人")) //暗送
			m.SetBody(message.contentType, message.body) // 正文

			//m.Attach("我是附件")  //添加附件
			if err := d.DialAndSend(m); err != nil {
				log.Println("发送失败", err)
				return
			} else {
				log.Println("done.发送成功")
			}
		}
	}()
	for i := 0; i < len(sendTo); i++ {
		<-done
	}

	return nil
}

func Cron()  {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		//message := &Message{
		//	email: "jikequan@163.com",
		//	subject: "第三期",
		//	body: "<p>我爱你中国</p>",
		//	contentType: "text/html",
		//}
		//Send(message)
		log.Println("cron running:", i)
	})
	c.AddFunc("@every 1h1m", func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
}