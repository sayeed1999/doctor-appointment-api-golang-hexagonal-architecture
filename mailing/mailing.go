package mailing

import (
	"crypto/tls"
	"log"

	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	gomail "gopkg.in/mail.v2"
)

func ListenForMail(mailChan chan domain.MailData, host string, port int, sender string, password string) {
	go func() {
		for {
			msg := <-mailChan
			sendMail(msg, host, port, sender, password)
		}
	}()
}

func sendMail(m domain.MailData, host string, port int, sender string, password string) {

	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", m.To)
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody("text/html", m.Content)

	// Settings for SMTP Server
	d := gomail.NewDialer(host, port, sender, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now Send E-mail
	if err := d.DialAndSend(msg); err != nil {
		log.Println("Error occured while sending email...")
		log.Println(err)
	} else {
		log.Println("email sent!")
	}
}
