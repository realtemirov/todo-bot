package storage

import (
	"time"

	"github.com/realtemirov/projects/tgbot/model"
)

type StorageI interface {
	User() UserI
	Todo() TodoI
}

type UserI interface {
	Add(user *model.User) (int64, error)
	Get(id int64) (*model.User, error)
	GetAll() ([]*model.User, error)
	SetAction(id int64, action string) error
	GetAction(id int64) (string, error)
}

type TodoI interface {
	Create(todo *model.Todo) (string, error)
	GetByID(id string) (*model.Todo, error)
	GetAllByUserID(id int64, done bool) ([]*model.Todo, error)
	GetAllNotificationTimes() ([]*model.Notification, error)
	GetAllDeadlineTimes() ([]*model.Deadline, error)

	GetNoSet(id int64) (*model.Todo, error)
	AddText(id int64, text string) error
	AddPhotoURL(id int64, photoURL string) error
	AddFileURL(id int64, fileURL string) error

	AddTime(id int64, date *time.Time) error
	AddHour(id int64, hour *time.Duration, column string) (*model.Todo, error)

	SetIsSet(id int64) error
	Done(id string, done bool) error
	DeleteByID(id string) error
	DeleteSetIsFalse(id int64) error
}
