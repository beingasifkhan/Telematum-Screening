package main

import (
	"log"
	"net/http"
	"screening/services"
)

func main() {
	services.SetupJsonApi()
	//checking for errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting server: ", err)
	}
}
