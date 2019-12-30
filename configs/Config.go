package configs

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func GetDialer() *gomail.Dialer {
	var (
		HOST = os.Getenv("SMTP_HOST")
		PORT, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
		USERNAME = os.Getenv("SMTP_USERNAME")
		PASSWORD = os.Getenv("SMTP_PASSWORD")
	)
	dialer := gomail.NewDialer(HOST, PORT, USERNAME, PASSWORD)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify:true}
	return dialer
}
