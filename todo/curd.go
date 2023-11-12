package todo

import (
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/abhinandpn/FirebaseTodo/model"
	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(clint *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. bind the 	recived/request json into a struct
		// 2. fill in a few fields
		// 3. save in to firestore
		var todo model.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		now := time.Now()
		todo.CreatedAt = now
		todo.UpdatedAt = now

		ref := clint.Collection(model.TODO_COLLECTION).NewDoc()

		_, err := ref.Set(c, map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"createdAt":   todo.CreatedAt,
			"updatedAt":   todo.UpdatedAt,
			"completed":   false,
		})

		if err != nil {
			log.Printf("an error has occured")
			c.JSON(http.StatusInternalServerError, "")
			return
		}
	}
}
