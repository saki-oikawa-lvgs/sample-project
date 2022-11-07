package service

import (
	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/common"
	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/customTypes"
)

func TodosGet() ([]*customTypes.Todo, error) {

	db := common.InitDb()
	// defer db.Close()

	var todos []*customTypes.Todo
	err := db.Table("todos").Find(&todos).Error

	return todos,err
}

func TodoCreate(input customTypes.NewTodo) (*customTypes.Todo, error) {
	
	db := common.InitDb()
	// defer db.Close()

	newTodo := customTypes.Todo{
		Text: input.Text,
	}

	err := db.Table("todos").Create(&newTodo).Error

	return &newTodo,err
}

func TodoFindByID(todoID int) (*customTypes.Todo, error){
	
	db := common.InitDb()
	// defer db.Close()

	var todo customTypes.Todo
	err := db.Table("todos").Where("id = ?", todoID).First(&todo).Error

	return &todo,err
}