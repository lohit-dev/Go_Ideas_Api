package main

import (
	"log"
	"test_project/test/cmd/app"
	_ "test_project/test/docs"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	log.Println("Server starting at http://localhost:8080")
	if err := app.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
