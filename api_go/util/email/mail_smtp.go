package email

import (
	"api/util/log"
	"net/smtp"
	"strconv"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
}

var (
	receivers []string
	subject string
	body string
)

// func New()*Config{
// 	return &Config{config.EmailConfig.Username, config.EmailConfig.Password, config.EmailConfig.Host, config.EmailConfig.Port}
// }

func (emailConf *Config)Receivers(r []string)*Config{
	receivers = r
	return emailConf
}

func (emailConf *Config)Subject(s string)*Config{
	subject = s
	return emailConf
}

func (emailConf *Config)Body(s string)*Config{
	body = s
	return emailConf
}

func (emailConf *Config)Send(){
	emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)
	sender := emailConf.Username
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"


	subject := "Subject: " + subject + "!\n"
	message := []byte(subject + mime + body)

	////message := []byte("To: recipient@example.net\r\n" +
	////	"Subject: discount Gophers!\r\n" +
	////	"\r\n" +
	////	"This is the email body. Hello from Go SMTP vendor.\r\n") // your message


	//// send out the email
	err := smtp.SendMail(emailConf.Host+":"+strconv.Itoa(emailConf.Port), //convert port number from int to string
		emailauth,
		sender,
		receivers,
		message,
	)

	log.Error(err)
}
