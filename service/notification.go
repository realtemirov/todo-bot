package service

import (
	"fmt"
	"time"

	"github.com/realtemirov/projects/tgbot/model"
	"github.com/realtemirov/projects/tgbot/storage"
)

type notifService struct {
	notificationRepo storage.NotificationI
}

func NewNotifRepository(notificationRepo storage.NotificationI) *notifService {
	return &notifService{
		notificationRepo: notificationRepo,
	}
}

func (t *notifService) Create(notif *model.Notification) (string, error) {
	fmt.Println("keldi", notif.ID)
	n := &model.Notification{
		Base: model.Base{
			ID:        notif.ID,
			CreatedAt: time.Now(),
		},
		Todo_ID:    notif.Todo_ID,
		User_ID:    notif.User_ID,
		Notif_date: notif.Notif_date,
	}
	fmt.Println("Hello from service")
	return t.notificationRepo.Create(n)

}

func (t *notifService) GetByUserId(id int64) ([]*model.Notification, error) {
	return t.notificationRepo.GetByUserId(id)
}

func (t *notifService) GetAll() ([]*model.Notification, error) {
	return t.notificationRepo.GetAll()
}
