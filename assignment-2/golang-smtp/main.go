//package main
//
//import (
//	"bytes"
//	"crypto/tls"
//	"fmt"
//	"log"
//	"net/smtp"
//	"strings"
//)
//
//func main() {
//	//dialing, establishing 2 way connection channel to SMTP server
//	c, err := smtp.Dial("smtp.gmail.com:465")
//	// in case connection fails, handle the error
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer c.Close()
//	// Set the sender and recipient.
//	c.Mail("indratejagmail.com")
//	c.Rcpt("binkkatal.r@gmail.com")
//	c.Rcpt("ravitejaraja.k.v@gmail.com")
//	// Send the email body.
//	wc, err := c.Data()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer wc.Close()
//	buf := bytes.NewBufferString("This is the email body.")
//	if _, err = buf.WriteTo(wc); err != nil {
//		log.Fatal(err)
//	}
//}

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	senderEmail string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderEmail)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func main() {
	mail := Mail{}
	mail.senderEmail = "indra.lf62@gmail.com"
	mail.toIds = []string{"ravitejaraja.k.v@gmail.com", "bindz.kanagala@gmail.com"}
	mail.subject = "Sent using Golang"
	mail.body = "Hello guys,\n\nAre we meeting today?"

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("", mail.senderEmail, "chintakayala1991", smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderEmail); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
