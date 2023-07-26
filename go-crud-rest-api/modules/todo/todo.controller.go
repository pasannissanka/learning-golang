package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/db"
	todoentity "github.com/pasannissanka/learning-golang/go-crud-rest-api/modules/todo/entity"
)

func GetTodos(c *gin.Context) {
	todos := []todoentity.Todo{}
	result := db.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error getting todos",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "get todos",
		"data":    todos,
	})
}

func CreateTodos(c *gin.Context) {
	todo := todoentity.Todo{Title: "L123", Description: "L123", Completed: false}
	result := db.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error creating todo",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "create todos",
		"data":    todo,
	})
}
