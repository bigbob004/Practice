package main

import (
	"log"
	web "webScraper/internal/router_and_server"
)

func main() {
	srv := web.GetHTTPServer()
	log.Print("http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

}
