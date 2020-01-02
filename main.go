package main

import (
	"QueueMail/controllers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := gotenv.Load(); if err != nil {
		panic(err)
	}
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/send", controllers.SendMail).Methods("POST")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
