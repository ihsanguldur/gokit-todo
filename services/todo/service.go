package todo

import "errors"

var (
	ErrUnexpectedDatabase = errors.New("unexpected database error")
)

type Service interface {
	CreateTodo(todo Todo) (string, error)
	ListTodo() ([]Todo, error)
	UpdateTodo(id uint, todo Todo) (string, error)
}
