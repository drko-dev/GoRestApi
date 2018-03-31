package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/route"
	"time"
)

func main() {

	var chain, key, port string
	if len(os.Args) > 2 {
		chain = os.Args[1]
		key = os.Args[2]
		port = os.Args[3]
	} else if len(os.Args) == 2 {
		chain = ""
		key = ""
		port = os.Args[1]
	} else {
		chain = "/etc/ssl/certs/fullchain.pem"
		key = "/etc/ssl/certs/privkey.pem"
		port = "1700"
	}

	router := route.NewMuxRouter()

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando puerto " + port + " ...")
	if chain == "" || key == "" {
		log.Println(server.ListenAndServe())
	} else {
		log.Println(server.ListenAndServeTLS(chain, key))
	}
}
