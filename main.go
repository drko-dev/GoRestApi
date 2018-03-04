package main

import (
	"log"
	"net/http"
	"rest-api/route"
	"time"
)

func main() {

	router := route.NewMuxRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Escuchando...")
	log.Fatal(server.ListenAndServe())

}
