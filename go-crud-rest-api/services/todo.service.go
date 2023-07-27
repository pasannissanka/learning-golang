package services

import (
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/db"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/entities"
	"gorm.io/gorm/clause"
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

func GetTodoById(id string) (entities.Todo, error) {
	todo := entities.Todo{}
	result := db.DB.First(&todo, id)

	if result.Error != nil {
		return entities.Todo{}, result.Error
	}

	return todo, nil
}

func UpdateTodo(id string, todo entities.Todo) (entities.Todo, error) {
	find := db.DB.First(&entities.Todo{}, id)
	if find.Error != nil {
		return entities.Todo{}, find.Error
	}
	var updatedTodo entities.Todo

	result := db.DB.Model(&updatedTodo).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(entities.Todo{Title: todo.Title, Description: todo.Description, Completed: todo.Completed})

	if result.Error != nil {
		return entities.Todo{}, result.Error
	}

	return updatedTodo, nil
}

func DeleteTodoById(id string) error {
	find := db.DB.First(&entities.Todo{}, id)
	if find.Error != nil {
		return find.Error
	}

	result := db.DB.Delete(&entities.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
