package main

import (
	"log"
	"net/http"

	"github.com/gopherskatowice/todolist-backend/server"
)

func main() {
	rt := server.RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8080", rt))
}
