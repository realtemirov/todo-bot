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
			Base:         model.Base{ID: uuid.New().String(), CreatedAt: time.Now()},
			User_ID:      todo.User_ID,
			Title:        todo.Title,
			Description:  "",
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

func (t *todoService) GetAll(id int64) ([]*model.Todo, error) {
	return t.todoRepo.GetAllByUserID(id)
}

func (t *todoService) GetAllNotificationTimes() ([]*model.Notification, error) {
	return t.todoRepo.GetAllNotificationTimes()
}

func (t *todoService) AddDescription(id int64, description string) error {
	return t.todoRepo.AddDescription(id, description)
}

func (t *todoService) AddNotification(id int64, deadline *time.Time) error {
	return t.todoRepo.AddNotification(id, deadline)
}

func (t *todoService) AddPhotoURL(id int64, photoURL string) error {
	return t.todoRepo.AddPhotoURL(id, photoURL)
}

func (t *todoService) AddFileURL(id int64, fileURL string) error {
	return t.todoRepo.AddFileURL(id, fileURL)
}

func (t *todoService) AddSetMonthToDeadLine(id int64, month *time.Time) error {
	return t.todoRepo.AddSetMonthToDeadLine(id, month)
}

func (t *todoService) AddSetDayToDeadLine(id int64, day *time.Time) error {
	return t.todoRepo.AddSetDayToDeadLine(id, day)
}

func (t *todoService) AddSetHourToDeadLine(id int64, hour *time.Time) error {
	return t.todoRepo.AddSetHourToDeadLine(id, hour)
}

func (t *todoService) SetIsSet(id int64) error {
	return t.todoRepo.SetIsSet(id)
}

func (t *todoService) Done(id string) error {
	return t.todoRepo.Done(id)
}

func (t *todoService) DeleteByID(id string) error {
	return t.todoRepo.DeleteByID(id)
}

func (t *todoService) DeleteSetIsFalse(id int64) error {
	return t.todoRepo.DeleteSetIsFalse(id)
}
