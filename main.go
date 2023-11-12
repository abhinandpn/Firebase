package main

import (
	"log"

	"github.com/abhinandpn/FirebaseTodo/firebase"
	"github.com/abhinandpn/FirebaseTodo/handlers"
	"github.com/abhinandpn/FirebaseTodo/todo"
	"github.com/gin-gonic/gin"
)

func main() {

	client, err := firebase.CreateClint()
	if err != nil {
		log.Printf("error while creating clint %v ", err)
		return
	}

	r := gin.Default()
	r.GET("/api/health", handlers.HealthcheckHandler())
	r.POST("/todo/new", todo.CreateTodoHandler(client))
	r.Run(":8081")
}
