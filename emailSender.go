package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func send(to []string) error {

	from := "testsender.genesis.education@gmail.com"
	password := "fpizvtmeaucfodap"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.New("").Parse(Message)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Current BTC to UAH exchange rate \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Rate string
	}{
		Rate: fmt.Sprintf("%f", getCurrentBTCToUAHRate()),
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
