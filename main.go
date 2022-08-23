package main

import (
	"github.com/andreabreu76/encurtadorV2/config"
	"log"
	"net/http"
)

func main() {
	session, err := config.ConnectMongo()
	if err != nil {
		log.Fatalf("Error connecting to mongo: %v", err)
	}

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Printf("Server started with %v", session)
}
