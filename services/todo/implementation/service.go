package implementation

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	todoSvc "go-kit-todo/services/todo"
)

type service struct {
	repository todoSvc.Repository
	logger     log.Logger
}

func New(repository todoSvc.Repository, logger log.Logger) todoSvc.Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s service) CreateTodo(todo todoSvc.Todo) (string, error) {
	logger := log.With(s.logger, "method", "CreateTodo")
	if err := s.repository.Create(todo); err != nil {
		_ = level.Error(logger).Log("err", err.Error())
		return "", todoSvc.ErrUnexpectedDatabase
	}

	return "todo created", nil
}

func (s service) ListTodo() ([]todoSvc.Todo, error) {
	logger := log.With(s.logger, "method", "ListTodo")
	todos, err := s.repository.List()
	if err != nil {
		_ = level.Error(logger).Log("err", err.Error())
		return nil, todoSvc.ErrUnexpectedDatabase
	}

	return todos, nil
}

func (s service) UpdateTodo(id uint, todo todoSvc.Todo) (string, error) {
	logger := log.With(s.logger, "method", "UpdateTodo")
	if err := s.repository.Update(id, todo); err != nil {
		_ = level.Error(logger).Log("err", err.Error())
		return "", todoSvc.ErrUnexpectedDatabase
	}

	return "todo updated", nil
}
