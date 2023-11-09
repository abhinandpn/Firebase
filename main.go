package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializong the app : %v", err)
	}

	client, err := app.Firestore(context.TODO())
	if err != nil {
		log.Fatal("error geting firestore clint %v ", err)
	}

	collRef := client.Collection("Firebase-Go")
	docref := collRef.Doc("users")
}
