package controllers

import (
	"QueueMail/configs"
	"QueueMail/structs"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func SendMail(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[+] Decoding Data")
	decoder := json.NewDecoder(request.Body)
	var queue structs.Queue
	err := decoder.Decode(&queue);
	if err != nil {
		fmt.Println("ERR : " + err.Error())
	}
	go SendingMailAsync(queue)
	fmt.Println("======== Data ========")
	fmt.Println("[*] Subject : " + queue.Email.Subject)
	fmt.Println("[*] Target : " + strings.Join(queue.Addresses[:], ";"))
	fmt.Println("======================")
}

func SendingMailAsync(queue structs.Queue) {
	fmt.Println("[+] Initiate Mailing")
	dialer := configs.GetDialer()
	var targets []structs.Target
	fmt.Println("[+] Sending Email")
	for _, address := range queue.Addresses {
		fmt.Println("[-] Send email to ", address)
		email := gomail.NewMessage()
		email.SetHeader("From", os.Getenv("SENDER_EMAIL"))
		email.SetHeader("Subject", queue.Email.Subject)
		email.SetHeader("To", address)
		email.SetBody("text/html", queue.Email.Content)
		err := dialer.DialAndSend(email);
		if err != nil {
			targets = append(targets, structs.Target{
				Address: address,
				Status:  "failed",
				Message: err.Error(),
			})
		} else {
			targets = append(targets, structs.Target{
				Address: address,
				Status:  "success",
				Message: "Mail sent successfully",
			})
		}
	}
	fmt.Println("[+] Create Log")
	createLog(queue, targets)
	fmt.Println("[+] Sending Done")
}

func createLog(queue structs.Queue, targets []structs.Target) {
	task := structs.Task{
		Email:   queue,
		Targets: targets,
		Date:    time.Now().Format("2006-01-02 15:04:05"),
	}

	client, coll, err := configs.GetClientCollection();
	if err != nil {
		log.Fatal(err)
		return
	}

	defer client.Disconnect(context.Background())

	data, err := coll.InsertOne(context.TODO(), task);
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data Inserted : ", data.InsertedID)
}
