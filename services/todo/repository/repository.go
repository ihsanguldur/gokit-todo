package repository

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	todoSvc "go-kit-todo/services/todo"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func New(db *gorm.DB, logger log.Logger) todoSvc.Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "repository"),
	}
}

func (r repository) Create(todo todoSvc.Todo) error {
	if err := r.db.Create(&todo).Error; err != nil {
		_ = level.Error(r.logger).Log("err", err.Error())
		return err
	}

	return nil
}

func (r repository) List() ([]todoSvc.Todo, error) {
	var todos []todoSvc.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		_ = level.Error(r.logger).Log("err", err.Error())
		return nil, err
	}

	return todos, nil
}

func (r repository) Update(id uint, todo todoSvc.Todo) error {
	if err := r.db.Model(&todoSvc.Todo{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		_ = level.Error(r.logger).Log("err", err.Error())
		return err
	}

	return nil
}
