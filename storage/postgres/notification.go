package postgres

import "github.com/jmoiron/sqlx"

type notificationRepo struct {
	db *sqlx.DB
}

func NewNotificationRepo(db *sqlx.DB) *notificationRepo {
	return &notificationRepo{
		db: db,
	}
}

func (n *notificationRepo) Send() {

}
func (n *notificationRepo) Check() {

}
