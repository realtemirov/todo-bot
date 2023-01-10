package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/realtemirov/projects/tgbot/model"
	"github.com/realtemirov/projects/tgbot/storage"
)

type todoService struct {
	todoRepo storage.TodoI
}

func NewTodoRepository(todoRepo storage.TodoI) *todoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t *todoService) Create(todo *model.Todo) (string, error) {
	return t.todoRepo.Create(
		&model.Todo{
			Base:         model.Base{ID: uuid.New().String(), CreatedAt: time.Now().Add(time.Hour * 5)},
			User_ID:      todo.User_ID,
			Text:         todo.Text,
			Photo_URL:    "",
			File_URL:     "",
			Deadline:     &time.Time{},
			Is_Set:       false,
			Notification: &time.Time{},
		},
	)
}

func (t *todoService) GetByID(id string) (*model.Todo, error) {
	return t.todoRepo.GetByID(id)
}

func (t *todoService) GetAllByUserID(id int64, done bool) ([]*model.Todo, error) {
	return t.todoRepo.GetAllByUserID(id, done)
}

func (t *todoService) GetAllNotificationTimes() ([]*model.Notification, error) {
	return t.todoRepo.GetAllNotificationTimes()
}

func (t *todoService) GetNoSet(id int64) (*model.Todo, error) {
	return t.todoRepo.GetNoSet(id)
}

func (t *todoService) AddText(id int64, text string) error {
	return t.todoRepo.AddText(id, text)
}

func (t *todoService) AddPhotoURL(id int64, photoURL string) error {
	return t.todoRepo.AddPhotoURL(id, photoURL)
}

func (t *todoService) AddFileURL(id int64, fileURL string) error {
	return t.todoRepo.AddFileURL(id, fileURL)
}

func (t *todoService) AddTime(id int64, date *time.Time) error {
	return t.todoRepo.AddTime(id, date)
}

func (t *todoService) AddHour(id int64, hour *time.Duration, column string) (*model.Todo, error) {
	return t.todoRepo.AddHour(id, hour, column)
}

func (t *todoService) SetIsSet(id int64) error {
	return t.todoRepo.SetIsSet(id)
}

func (t *todoService) Done(id string, done bool) error {
	return t.todoRepo.Done(id, done)
}

func (t *todoService) DeleteByID(id string) error {
	return t.todoRepo.DeleteByID(id)
}

func (t *todoService) DeleteSetIsFalse(id int64) error {
	return t.todoRepo.DeleteSetIsFalse(id)
}
