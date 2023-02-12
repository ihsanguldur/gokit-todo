package middleware

import (
	"github.com/go-kit/log"
	todoSvc "go-kit-todo/services/todo"
	"time"
)

type loggerMiddleware struct {
	logger log.Logger
	next   todoSvc.Service
}

func New(logger log.Logger) Middleware {
	return func(svc todoSvc.Service) todoSvc.Service {
		return &loggerMiddleware{
			logger: logger,
			next:   svc,
		}
	}
}

func (l *loggerMiddleware) CreateTodo(todo todoSvc.Todo) (msg string, err error) {
	defer func(begin time.Time) {
		_ = l.logger.Log("method", "CreateTodo", "TodoID", todo.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return l.next.CreateTodo(todo)
}

func (l *loggerMiddleware) ListTodo() (todos []todoSvc.Todo, err error) {
	defer func(begin time.Time) {
		_ = l.logger.Log("method", "ListTodo", "took", time.Since(begin), "err", err)
	}(time.Now())
	return l.next.ListTodo()
}

func (l *loggerMiddleware) UpdateTodo(id uint, todo todoSvc.Todo) (msg string, err error) {
	defer func(begin time.Time) {
		_ = l.logger.Log("method", "UpdateTodo", "todoID", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return l.next.UpdateTodo(id, todo)
}
