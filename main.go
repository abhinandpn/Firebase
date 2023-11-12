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
	r.POST("/todo/new", todo.CreateTodoHandler(client)) // create
	r.GET("/todo/full", todo.ListTodos(client))         // get all
	r.GET("/todo/:id", todo.GetTodoById(client))        // get by id
	r.PATCH("/todo/:id", todo.UpdateTodo(client))       // update todo
	r.DELETE("/todo/:id", todo.DeleteTodo(client))      // delete todo

	r.Run(":8081")
}
