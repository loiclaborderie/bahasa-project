package main

import (
	"log"

	"github.com/loiclaborderie/bahasa-project/routes"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server starting on :8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
