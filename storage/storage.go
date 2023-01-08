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
	GetAllByUserID(id int64) ([]*model.Todo, error)
	GetAllNotificationTimes() ([]*model.Notification, error)

	//GetNoSet(id int64) (*model.Todo, error)
	AddDescription(id int64, description string) error
	AddNotification(id int64, deadline *time.Time) error
	AddPhotoURL(id int64, photoURL string) error
	AddFileURL(id int64, fileURL string) error

	AddSetMonthToDeadLine(id int64, month *time.Time) error
	AddSetDayToDeadLine(id int64, day *time.Time) error
	AddSetHourToDeadLine(id int64, hour *time.Time) error

	SetIsSet(id int64) error
	Done(id string) error
	DeleteByID(id string) error
	DeleteSetIsFalse(id int64) error
}
