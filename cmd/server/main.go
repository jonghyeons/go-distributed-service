package main

import (
	"github.com/jonghyeons/go-distributed-service/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
