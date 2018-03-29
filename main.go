package main

import (
	"log"
	"net/http"
	"rest-api/route"
	"time"
)

func main() {

	/* 	var chain, key, port string
	   	if len(os.Args) > 1 {
	   		chain = os.Args[1]
	   		key = os.Args[2]
	   		port = os.Args[3]
	   	} else {
	   		chain = "/etc/ssl/certs/fullchain.pem"
	   		key = "/etc/ssl/certs/privkey.pem"
	   		port = "1700"
	   	} */

	router := route.NewMuxRouter()

	server := &http.Server{
		Addr:           ":1700",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando puerto 1700 ...")
	log.Println(server.ListenAndServe())
}
