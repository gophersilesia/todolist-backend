package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gopherskatowice/todolist-backend/server"
)

func main() {
	rt := server.RegisterHandlers()
	fmt.Println("Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}
