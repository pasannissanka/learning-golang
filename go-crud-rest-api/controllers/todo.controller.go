package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/entities"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/services"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/utils"
)

func TodoController(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
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

	router.POST("/", func(c *gin.Context) {
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

	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		todo, err := services.GetTodoById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "found todo",
			"data":    todo,
		})
	})

	router.POST("/:id", func(c *gin.Context) {
		id := c.Param("id")
		body := entities.Todo{}

		if err := c.ShouldBindJSON(&body); err != nil {
			errors := utils.ExtractValidationErrors(err.(validator.ValidationErrors))
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors, "message": "validation error"})
			return
		}

		todo, err := services.UpdateTodo(id, body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "todo updated todo",
			"data":    todo,
		})
	})

	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := services.DeleteTodoById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "todo deleted",
		})
	})
}
