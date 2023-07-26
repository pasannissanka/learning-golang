package services

import (
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/db"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/entities"
)

func GetTodos() ([]entities.Todo, error) {
	todos := []entities.Todo{}
	result := db.DB.Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func CreateTodo(todo entities.Todo) (entities.Todo, error) {
	result := db.DB.Create(&todo)

	if result.Error != nil {
		return entities.Todo{}, result.Error
	}

	return todo, nil
}
