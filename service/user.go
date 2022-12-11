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
		Username:  user.Username,
		FirsName:  user.FirsName,
		LastName:  user.LastName,
		IsAdmin:   false,
		Photo_URL: user.Photo_URL,
		CreatedAt: time.Now(),
	})
}

// get
func (u *userService) Update(user *model.User) (*model.User, error) {
	return u.userRepo.Update(user)
}

// get by id
func (u *userService) GetByID(id int64) (*model.User, error) {
	return u.userRepo.Get(id)
}

func (u *userService) GetAll() ([]*model.User, error) {
	return u.userRepo.GetAll()
}

func (u *userService) AddAdmin(id int64) error {
	return u.userRepo.AddAdmin(id)
}

func (u *userService) RemoveAdmin(id int64) error {
	return u.userRepo.RemoveAdmin(id)
}
