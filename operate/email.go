package operate

import (
	"net/smtp"
	"strings"
)

func SendToMail(user, password, host, rssMsg string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + user + "\r\nFrom: " + user + "<" + user + ">" + "\r\nSubject: " + "MetaRssMsg" + "\r\n" + contentType + "\r\n\r\n" + rssMsg)
	sendTo := strings.Split(user, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
