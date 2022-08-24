package main

import (
	"github.com/andreabreu76/encurtadorV2/config"
	"github.com/andreabreu76/encurtadorV2/route"
	"log"
	"net/http"
)

func main() {

	router := route.SetupRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Printf("Server listening on %v", config.HttpPort)
}
