package structs

import (
	"encoding/json"
	"log"
)

type Log struct {
	Id      interface{} `json:"_id" bson:"_id"`
	Email   Queue       `json:"email"`
	Targets []Target    `json:"targets"`
	Date    string      `json:"date"`
}

type Task struct {
	Email   Queue    `json:"email"`
	Targets []Target `json:"targets"`
	Date    string   `json:"date"`
}

type Target struct {
	Address string `json:"address"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (task Task) String() string {
	data, err := json.Marshal(task);
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
