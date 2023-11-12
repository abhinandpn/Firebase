package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
)

func main() {

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	ProjectId := os.Getenv("ProjectId")
	
	// Firebase connection testing
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: ProjectId}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Printf("error initializing app : %v ", err)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Print("error when creating clint", err)
		log.Fatal(err)
	}

	ref := client.Collection("todos").NewDoc()

	result, err := ref.Set(ctx, map[string]interface{}{
		"title":       "A random todo ",
		"description": "learn golang",
	})

	if err != nil {
		log.Printf("error while creating a todo list : %v ", err)
	}
	log.Printf("Result is [ %v ]", result)
}
