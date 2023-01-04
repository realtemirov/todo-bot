package storage

import "github.com/realtemirov/projects/tgbot/model"

type StorageI interface {
	User() UserI
	Todo() TodoI
}

type UserI interface {
	Add(user *model.User) (int64, error)
	Get(id int64) (*model.User, error)
}

type TodoI interface {
	Create(todo *model.Todo) (string, error)
	GetByID(id string) (*model.Todo, error)
	GetAllByUserID(id int64) ([]*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Done(id string) error
	Delete(id string) error
}
