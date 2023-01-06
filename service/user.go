package service

import (
	"time"

	"github.com/realtemirov/projects/tgbot/model"
	"github.com/realtemirov/projects/tgbot/storage"
)

type userService struct {
	userRepo storage.UserI
}

func NewUserRepository(userRepo storage.UserI) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Add(user *model.User) (int64, error) {
	return u.userRepo.Add(&model.User{
		ID:        user.ID,
		CreatedAt: time.Now().Add(time.Hour * 5),
		Action:    user.Action,
	})
}

func (u *userService) GetByID(id int64) (*model.User, error) {
	return u.userRepo.Get(id)
}

func (u *userService) SetAction(id int64, action string) error {
	return u.userRepo.SetAction(id, action)
}

func (u *userService) GetAction(id int64) (string, error) {
	return u.userRepo.GetAction(id)
}

/*
Add(user *model.User) (int64, error)
Get(id int64) (*model.User, error)
SetAction(id int64, action string) error
GetAction(id int64) (string, error)
*/
