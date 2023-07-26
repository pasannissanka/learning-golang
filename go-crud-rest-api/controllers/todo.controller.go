package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/entities"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/services"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/utils"
)

func TodoController(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		todos, err := services.GetTodos()

		if err != nil {
			c.JSON(400, gin.H{
				"message": "error getting todos",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "get todos",
			"data":    todos,
		})
	})

	rg.POST("/", func(c *gin.Context) {
		body := entities.Todo{}

		if err := c.ShouldBindJSON(&body); err != nil {
			errors := utils.ExtractValidationErrors(err.(validator.ValidationErrors))
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors, "message": "validation error"})
			return
		}

		todo, err := services.CreateTodo(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "validation error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "create todos",
			"data":    todo,
		})
	})
}
