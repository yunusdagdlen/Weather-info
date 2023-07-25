package main

import (
	_ "encoding/json"
	_ "github.com/lib/pq"
	"log"
	"main/application"
	"net/http"
)

func main() {
	application.InitializeEndpoints()
	http.ListenAndServe(":8080", nil)
	log.Println("Server starded on: http.//localhost:8080")
}
