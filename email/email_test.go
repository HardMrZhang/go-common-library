package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"sync"
	"testing"
	"time"
)

/**
电子邮件在网络中传输和网页一样需要遵从特定的协议，常用的电子邮件协议包括 SMTP，POP3，IMAP。
其中邮件的创建和发送只需要用到 SMTP协议，所以本文也只会涉及到SMTP协议。
SMTP 是 Simple Mail Transfer Protocol 的简称，即简单邮件传输协议。
*/
func Test_Email_01(t *testing.T) {
	/**
	发件人，收件人，密件抄送和抄送字段
	文字和HTML邮件正文
	附件
	阅读收据
	自定义标题
	*/
	e := email.NewEmail()
	e.From = "zzy <16616236933@163.com>"
	e.To = []string{"571135673@qq.com"}
	e.Subject = "send cjq"
	e.Text = []byte("this is a email from zzy!")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "18816236933@163.com", "GHRHLAUJHLEGLZTM", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}

//实现邮件抄送
func Test_Email_02(t *testing.T) {
	e := email.NewEmail()
	e.From = "zzy <18816236933@163.com>"
	e.To = []string{"1121562683@qq.com"}
	//设置抄送如果抄送多人逗号隔开
	e.Cc = []string{"571135673@qq.com"}
	//设置秘密抄送
	e.Bcc = []string{"571135673@qq.com"}
	e.Subject = "send cjq"
	e.Text = []byte("this is a email from zzy!")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "18816236933@163.com", "GHRHLAUJHLEGLZTM", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}

//邮件发送html
func Test_Email_03(t *testing.T) {
	e := email.NewEmail()
	e.From = "zzy <18816236933@163.com>"
	e.To = []string{"571135673@qq.com"}
	e.Subject = "send cjq"
	e.HTML = []byte(`<li><a href="http://www.baidu.com.cn">百度</a></li>`)
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "18816236933@163.com", "GHRHLAUJHLEGLZTM", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}

//发送附件
func Test_Email_04(t *testing.T) {
	e := email.NewEmail()
	e.From = "zzy <18816236933@163.com>"
	e.To = []string{"571135673@qq.com"}
	e.Subject = "send cjq"
	e.AttachFile("./go.mod")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "18816236933@163.com", "GHRHLAUJHLEGLZTM", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}

//每一次调用send都会和smtp建立一次连接,如果频繁的发送邮件,就会频繁的建立连接,这个时候用连接数,复用网络连接】
func Test_Email_05(t *testing.T) {
	ch := make(chan *email.Email, 10)
	p, err := email.NewPool("smtp.163.com:25", 4, smtp.PlainAuth("", "18816236933@163.com", "GHRHLAUJHLEGLZTM", "smtp.163.com"))
	if err != nil {
		log.Fatal("建立连接池失败--", err)
	}
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email:%v send error:%v\n", e, err)
				}
			}
		}()
	}
	for i := 0; i < 4; i++ {
		e := email.NewEmail()
		e.From = "zzy <18816236933@163.com>"
		e.To = []string{"571135673@qq.com"}
		e.Subject = "send cjq"
		e.Text = []byte(fmt.Sprintf("this is %d email from zzy!", i+1))

		ch <- e
	}
	close(ch)
	wg.Wait()
}
