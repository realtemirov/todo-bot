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
		CreatedAt: time.Now(),
	})
}

func (u *userService) GetByID(id int64) (*model.User, error) {
	return u.userRepo.Get(id)
}
