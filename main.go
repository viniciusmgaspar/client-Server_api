package main

import (
	"log"
	"net/http"

	"github.com/viniciusmgaspar/client-Server_api/client"
	"github.com/viniciusmgaspar/client-Server_api/database"
	"github.com/viniciusmgaspar/client-Server_api/server"
)

func main() {
	db, err := database.DBconnect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	go http.ListenAndServe(":8080", nil)
	server.Server(db)
	client.Execute()
}
