package todo

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Content string `json:"content"`
	State   bool   `json:"state" gorm:"default:false"`
	UserID  uint   `json:"user_id"`
}

type Repository interface {
	Create(todo Todo) error
	List() ([]Todo, error)
	Update(id uint, todo Todo) error
}
