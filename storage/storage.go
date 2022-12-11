package storage

import "github.com/realtemirov/projects/tgbot/model"

type StorageI interface {
	User() UserI
	Challenge() ChallengeI
	Todo() TodoI
	Notification() NotificationI
}

type UserI interface {
	Add(user *model.User) (int64, error)
	Get(id int64) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	GetAll() ([]*model.User, error)
	AddAdmin(id int64) error
	RemoveAdmin(id int64) error
}

type ChallengeI interface {
	Create()
	GetByID()
	GetAll()
	Update()
	Delete()
}
type TodoI interface {
	Create(todo *model.Todo) (string, error)
	GetByID(id string) (*model.Todo, error)
	GetAllByUserID(id int64) ([]*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Done(id string) error
	Delete(id string) error
}

type NotificationI interface {
	Send()
	Check()
}
