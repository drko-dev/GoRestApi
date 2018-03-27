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
		Addr:           ":1700",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Escuchando puerto 1700 ...")
	log.Fatal(server.ListenAndServeTLS("/etc/letsencrypt/live/visio.ctssysltd.com/fullchain.pem", " /etc/letsencrypt/live/visio.ctssysltd.com/privkey.pem"))

}
