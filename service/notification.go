package service

import (
	"time"

	"github.com/google/uuid"
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

	n := &model.Notification{
		Base: model.Base{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
		},
		Todo_ID:    notif.Todo_ID,
		User_ID:    notif.User_ID,
		Notif_date: notif.Notif_date,
	}

	return t.notificationRepo.Create(n)

}
