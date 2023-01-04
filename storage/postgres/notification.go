package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/projects/tgbot/model"
)

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

func (n *notificationRepo) Create(notif *model.Notification) (string, error) {

	var id string

	q := `INSERT INTO notification(id, todo_id, user_id, created_at, updated_at, deleted_at, notif_date)
		 VALUES( $1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	row := n.db.QueryRow(q, notif.Base.ID, notif.Todo_ID, notif.User_ID, notif.CreatedAt, notif.UpdatedAt, notif.DeletedAt, notif.Notif_date)
	err := row.Scan(&id)

	if err != nil {
		return "", err
	}
	fmt.Println("Repo", id)
	return id, nil
}

func (n *notificationRepo) GetByUserId(id int64) ([]*model.Notification, error) {

	res := make([]*model.Notification, 0)
	q := `SELECT id, todo_id, user_id, created_at, updated_at, deleted_at, notif_date from notification WHERE deleted_at = NULL`

	rows, err := n.db.Query(q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var notif model.Notification

		err = rows.Scan(&notif.ID, &notif.Todo_ID, &notif.User_ID, &notif.CreatedAt, &notif.UpdatedAt, &notif.DeletedAt, &notif.Notif_date)

		if err != nil {
			return nil, err
		}

		res = append(res, &notif)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (n *notificationRepo) GetAll() ([]*model.Notification, error) {
	q := `SELECT user_id, todo_id, notif_date FROM notification`
	rows, err := n.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Notification, 0)

	for rows.Next() {
		var notif model.Notification

		err = rows.Scan(&notif.User_ID, &notif.Todo_ID, &notif.Notif_date)
		if err != nil {
			return nil, err
		}

		res = append(res, &notif)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
