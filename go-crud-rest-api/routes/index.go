package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/controllers"
)

func GetRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		// todo routes
		todos := v1.Group("/todo")
		controllers.TodoController(todos)
	}
}
