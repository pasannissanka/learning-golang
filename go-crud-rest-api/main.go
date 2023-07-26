package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/db"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/routes"
)

func main() {
	r := gin.Default()
	db.Init()

	api := r.Group("/api")
	{
		routes.GetRoutes(api)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
