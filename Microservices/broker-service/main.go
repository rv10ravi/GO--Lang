package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {
	log.Println("Starting broker-service...")

	srv:= &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.Routes(),
	}

	err:= srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}