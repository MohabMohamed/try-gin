package models

func (todo Todo) Create() error {
	return DB.Create(&todo).Error
}

func (todo Todo) GetByID(id uint) error {
	return DB.First(&todo, id).Error
}

func (todo Todo) Delete() error {
	return DB.Delete(&todo).Error
}

func (todo Todo) UpdateByID() error {
	var oldTodo Todo
	var err error
	err = DB.First(&oldTodo, todo.ID).Error
	if err != nil {
		return err
	}
	return DB.Save(&todo).Error
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	err := DB.Find(&todos).Error
	return todos, err
}
