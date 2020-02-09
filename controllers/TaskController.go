package controllers

import (
	"QueueMail/configs"
	"QueueMail/structs"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Tasks(writer http.ResponseWriter, request *http.Request) {
	filter := bson.D{}
	if request.URL.Query()["address"] != nil{
		filter = append(filter, bson.E{"email.addresses", primitive.Regex{
			Pattern: request.URL.Query()["address"][0],
			Options: "i",
		}})
	}

	if request.URL.Query()["status"] != nil{
		filter = append(filter, bson.E{"targets.status", primitive.Regex{
			Pattern: request.URL.Query()["status"][0],
			Options: "i",
		}})
	}

	if request.URL.Query()["email"] != nil{
		filter = append(filter, bson.E{"email.subject", primitive.Regex{
			Pattern: request.URL.Query()["email"][0],
			Options: "i",
		}})

		filter = append(filter, bson.E{"email.content", primitive.Regex{
			Pattern: request.URL.Query()["email"][0],
			Options: "i",
		}})
	}

	if request.URL.Query()["status_message"] != nil{
		filter = append(filter, bson.E{"targets.message", primitive.Regex{
			Pattern: request.URL.Query()["status_message"][0],
			Options: "i",
		}})
	}

	if request.URL.Query()["date"] != nil{
		filter = append(filter, bson.E{"date", primitive.Regex{
			Pattern: request.URL.Query()["date"][0],
			Options: "i",
		}})
	}

	var errors []error
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	client, coll, err := configs.GetClientCollection();
	if err != nil {
		errors = append(errors, err)
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	var tasks []*structs.Log
	res, err := coll.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"date", -1}}));
	if err != nil {
		errors = append(errors, err)
		log.Fatal(err)
	}

	for res.Next(context.TODO()) {
		var task structs.Log
		err := res.Decode(&task);
		if err != nil {
			errors = append(errors, err)
			log.Fatal(err)
		}
		tasks = append(tasks, &task)
	}

	data := ApiResponse{
		Status:  "success",
		Message: "",
		Data:    tasks,
	}

	responseData, err := json.Marshal(data)

	writer.Write(responseData)
}

func Task(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	var task structs.Log
	client, coll, err := configs.GetClientCollection();
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	var data ApiResponse

	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)

	if err != nil {
		data = ApiResponse{
			Status:  "failed",
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		data = ApiResponse{
			Status:  "success",
			Message: "Data Found",
			Data:    task,
		}
	}

	responseData, err := json.Marshal(data); if err != nil {
		data = ApiResponse{
			Status:  "failed",
			Message: err.Error(),
			Data:    nil,
		}
		responseData, _ = json.Marshal(data)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseData)
}

func Resend(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	var task structs.Log
	var data ApiResponse
	client, coll, err := configs.GetClientCollection();
	if err != nil {
		data = ApiResponse{
			Status:  "failed",
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		defer client.Disconnect(context.Background())

		err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
		if err != nil {
			data = ApiResponse{
				Status:  "failed",
				Message: err.Error(),
				Data:    nil,
			}
		} else {
			data = ApiResponse{
				Status:  "success",
				Message: "Sending Email",
				Data:    nil,
			}
		}
	}

	responseData, err := json.Marshal(data); if err != nil {
		data = ApiResponse{
			Status:  "failed",
			Message: err.Error(),
			Data:    nil,
		}
		responseData, _ = json.Marshal(data)
	}

	writer.Write(responseData)

	SendingMailAsync(task.Email)
}
