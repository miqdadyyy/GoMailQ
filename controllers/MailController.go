package controllers

import (
	"QueueMail/configs"
	"QueueMail/structs"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
	"strings"
)

func SendMail(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[+] Decoding Data")
	decoder := json.NewDecoder(request.Body)
	var queue structs.Queue
	err := decoder.Decode(&queue); if err != nil {
		fmt.Println("ERR : " + err.Error())
	}
	go sendingMailAsync(queue)
	fmt.Println("======== Data ========")
	fmt.Println("[*] Subject : " + queue.Email.Subject)
	fmt.Println("[*] Target : " + strings.Join(queue.Addresses[:], ";"))
	fmt.Println("======================")
	PrintLog(request)
}

func sendingMailAsync(queue structs.Queue){
	fmt.Println("[+] Initiate Mailing")
	dialer := configs.GetDialer()
	fmt.Println("[+] Sending Email")
	email := gomail.NewMessage()
	email.SetHeader("From", os.Getenv("SENDER_EMAIL"))
	email.SetHeader("Subject", queue.Email.Subject)
	email.SetHeader("To", queue.Addresses...)
	email.SetBody("text/html", queue.Email.Content)
	err := dialer.DialAndSend(email); if err != nil {
		panic(err)
	}
	fmt.Println("[+] Sending Done")
}
