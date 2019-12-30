package controllers

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(`{"message": "This is main page of mailer"}`))

	PrintLog(request)
}

func PrintLog(r *http.Request)  {
	fmt.Println("HTTP : " + r.Method + " => " + r.URL.String())
}

