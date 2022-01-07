package main

import (
	"github.com/travisjeffery/proglog/tree/master/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
