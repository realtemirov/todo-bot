package service

import (
	"github.com/realtemirov/projects/tgbot/storage"
)

type Service struct {
	strg         storage.StorageI
	UserService  *userService
	TodoService  *todoService
	NotifService *notifService
}

func NewService(strg storage.StorageI) *Service {
	return &Service{
		strg:         strg,
		UserService:  NewUserRepository(strg.User()),
		TodoService:  NewTodoRepository(strg.Todo()),
		NotifService: NewNotifRepository(strg.Notification()),
	}
}
