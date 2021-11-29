package email

import (
	"api/util/log"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
)

type LoginAuth struct {
	username, password string
}

func NewLoginAuth(username, password string) smtp.Auth {
	return &LoginAuth{username, password}
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, nil
		}
	}
	return nil, nil
}

//信件內容設定
func sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte, subject string) {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		log.Error(err)
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		configs := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(configs); err != nil {
			log.Error(err)
		}
	}

	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				log.Error(err)
			}
		}
	}

	if err = c.Mail(from); err != nil {
		log.Error(err)
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			log.Error(err)
		}
	}
	w, err := c.Data()
	if err != nil {
		log.Error(err)
	}

	header := make(map[string]string)
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(msg)
	_, err = w.Write([]byte(message))

	if err != nil {
		log.Error(err)
	}
	err = w.Close()
	if err != nil {
		log.Error(err)
	}
	log.Error(c.Quit())
}

//外面呼叫寄送Email的Function
func SendMail(email string,subject string, message string) {
	auth := NewLoginAuth("hugo.c.yan@7officer.com", "730830wdst")
	//to := getAdminMail()
	msg := []byte(message)
	sendMail("smtp.gmail.com:587", auth, "hugo.c.yan@7officer.com", []string{email}, msg, subject)
}

////取得需要被寄送的 管理者Email (Table : admin_mail)
//func getAdminMail() (emailAry []string) {
//	mysql.Option(func(db *sql.DB) {
//		rows, err := db.Query("select email from p2p.admin_mail")
//		defer rows.Close()
//		if err == nil {
//			var email string
//			for rows.Next() {
//				err := rows.Scan(&email)
//				if err == nil {
//					emailAry = append(emailAry, email)
//				}
//			}
//		}
//	})
//	return emailAry
//}
