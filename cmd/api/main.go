package main

import (
	"log"

	"github.com/ranaincode/gpi-backend/internal/server"
)

func main() {
	srv := server.New()
	log.Println("GPI backend starting on : 8080")
	log.Fatal(srv.Start(":8080"))
}

