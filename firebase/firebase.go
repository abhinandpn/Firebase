package firebase

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/abhinandpn/FirebaseTodo/env"
)

func CreateClint() (*firestore.Client, error) {

	projectid, err := env.LoadEnv()
	if err != nil {
		fmt.Println("error while load env")
	}
	// Firebase connection testing
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectid}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Printf("error initializing app : %v ", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Print("error when creating clint", err)
		log.Fatal(err)
	}
	return client, err
}