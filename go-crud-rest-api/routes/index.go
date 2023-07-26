package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/modules/todo"
)

func GetRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		// todo routes
		todos := v1.Group("/todo")
		todos.GET("/", todo.GetTodos)
		todos.POST("/", todo.CreateTodos)
	}
}
