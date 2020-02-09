package main

import (
	"QueueMail/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router.Use(tokenMiddleware)
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/send", controllers.SendMail).Methods("POST")
	router.HandleFunc("/task", controllers.Tasks).Methods("GET")
	router.HandleFunc("/task/{id}", controllers.Task).Methods("GET")
	router.HandleFunc("/task/{id}/resend", controllers.Resend).Methods("GET")

	cors_ := cors.New(cors.Options{
		AllowedOrigins:         []string{
			os.Getenv("ALLOWED_ORIGINS"),
		},
		AllowedMethods:         []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
	})

	handler := cors_.Handler(router)

	http.Handle("/", router)
	fmt.Println("Listen on ", os.Getenv("BASEURL")+":"+os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("BASEURL") + ":" + os.Getenv("PORT"), handler))
}

func tokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		controllers.PrintLog(r)
		if r.Header.Get("key") == os.Getenv("KEY") {
			next.ServeHTTP(w, r)
		} else {
			w.Write([]byte(`{"message": "Key doesn't match'"}`))
		}
	})
}