package todo

import (
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/abhinandpn/FirebaseTodo/model"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func CreateTodoHandler(client *firestore.Client) func(c *gin.Context) {
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

		ref := client.Collection(model.TODO_COLLECTION).NewDoc()

		_, err := ref.Set(c, map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"createdAt":   todo.CreatedAt,
			"updatedAt":   todo.UpdatedAt,
			"completed":   false,
		})

		if err != nil {
			log.Printf("an error has occured")
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusCreated, todo)
	}
}

func ListTodos(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var result []model.Todo
		iter := client.Collection(model.TODO_COLLECTION).Documents(c)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, "")
				return
			}
			createdAtValue, exists := doc.Data()["createdAt"]
			if !exists || createdAtValue == nil {
				// Handle the case where createdAt is nil or doesn't exist
				continue
			}

			result = append(result, model.Todo{
				Id:          doc.Ref.ID,
				Title:       doc.Data()["title"].(string),
				Description: doc.Data()["description"].(string),
				CreatedAt:   createdAtValue.(time.Time),
				UpdatedAt:   doc.Data()["updatedAt"].(time.Time),
			})
			c.JSON(http.StatusOK, result)
		}
	}
}

func GetTodoById(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {

		id := c.Param("id")
		iter, err := client.Collection(model.TODO_COLLECTION).Doc(id).Get(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		doc := iter.Data()
		createdAtValue, exists := doc["createdAt"]
		if !exists || createdAtValue == nil {
			// Handle the case where createdAt is nil or doesn't exist
			c.JSON(http.StatusNotFound, "Todo not found")
			return
		}

		updatedAtValue, exists := doc["updatedAt"]
		if !exists || updatedAtValue == nil {
			// Handle the case where updatedAt is nil or doesn't exist
			c.JSON(http.StatusNotFound, "Todo not found")
			return
		}
		todo := &model.Todo{
			Id:          id,
			Title:       doc["title"].(string),
			Description: doc["description"].(string),
			CreatedAt:   createdAtValue.(time.Time),
			UpdatedAt:   updatedAtValue.(time.Time),
		}

		c.JSON(http.StatusOK, todo)

	}
}

func UpdateTodo(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {

		id := c.Param("id")
		var todo model.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo.Id = id
		todo.UpdatedAt = time.Now()
		_, err := client.
			Collection(model.TODO_COLLECTION).
			Doc(todo.Id).Set(c, map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"updatedAt":   todo.UpdatedAt,
			"completed":   todo.UpdatedAt,
		}, firestore.MergeAll)
		if err != nil {
			log.Printf("An error has occured : %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {

	}
}
