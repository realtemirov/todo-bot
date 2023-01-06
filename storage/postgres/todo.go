package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/projects/tgbot/model"
)

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) *todoRepo {
	return &todoRepo{
		db: db,
	}
}

// Create new TODO  => id, user_id, title, created_at, is_set
// Create(todo *model.Todo) (string, error)
func (t *todoRepo) Create(todo *model.Todo) (string, error) {

	var id string
	q := `INSERT INTO todos (id, created_at, user_id, title, description, photo_url, file_url, deadline, is_set, notification)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err := t.db.QueryRow(q, todo.ID, todo.CreatedAt, todo.User_ID, todo.Title, todo.Description, todo.Photo_URL, todo.File_URL, todo.Deadline, todo.Is_Set, todo.Notification).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Get todo by id  => user_id, title, description, photo_url, file_url
// GetByID(id string) (*model.Todo, error)
func (t *todoRepo) GetByID(id string) (*model.Todo, error) {
	fmt.Println("GetByID: ", id)
	q := `SELECT user_id, title, description, photo_url, file_url FROM todos WHERE id = $1`
	todo := model.Todo{}
	row := t.db.QueryRow(q, id)
	err := row.Scan(&todo.User_ID, &todo.Title, &todo.Description, &todo.Photo_URL, &todo.File_URL)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// Get todo by user_id  => id, title, description, photo_url, file_url
// GetAllByUserID(id int64) ([]*model.Todo, error)
func (t *todoRepo) GetAllByUserID(id int64) ([]*model.Todo, error) {
	q := `SELECT id, title, description, photo_url, file_url, deadline, FROM todos WHERE user_id = $1 AND deleted_at IS NULL and is_set = true ORDER BY deadline ASC`
	todos := []*model.Todo{}
	err := t.db.Select(&todos, q, id)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepo) GetAllNotificationTimes() ([]*model.Notification, error) {
	q := `SELECT id, notification FROM todos WHERE deleted_at IS NULL AND is_set = true`
	todos := []*model.Notification{}

	err := t.db.Select(&todos, q)

	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Add description to todo which is not set
// AddDescription(id string, description string) error
func (t *todoRepo) AddDescription(id int64, description string) error {
	q := `UPDATE todos SET description = $1 WHERE user_id = $2 AND is_set = false`
	_, err := t.db.Exec(q, description, id)
	if err != nil {
		return err
	}
	return nil
}

// Create new notification time stamp and add to todo which is not set
// AddNotification(id int64, deadline int64) error
func (t *todoRepo) AddNotification(id int64, notification *time.Time) error {
	q := `UPDATE todos SET notification = $1 WHERE user_id = $2 AND is_set = false`
	_, err := t.db.Exec(q, notification, id)
	if err != nil {
		return err
	}
	return nil
}

// Add photo url to todo which is not set
// AddPhotoURL(id int64, photoURL string) error
func (t *todoRepo) AddPhotoURL(id int64, url string) error {
	q := `UPDATE todos SET photo_url = $1 WHERE user_id = $2 AND is_set = false`
	_, err := t.db.Exec(q, url, id)
	if err != nil {
		return err
	}
	return nil
}

// Add file url to todo which is not set
// AddFileURL(id int64, fileURL string) error
func (t *todoRepo) AddFileURL(id int64, url string) error {
	q := `UPDATE todos SET file_url = $1 WHERE user_id = $2 AND is_set = false`
	_, err := t.db.Exec(q, url, id)
	if err != nil {
		return err
	}
	return nil
}

// Add set month to deadline of todo which is_set is false
// AddSetMonthToDeadLine(id int64, month *time.Time) error
func (t *todoRepo) AddSetMonthToDeadLine(id int64, month *time.Time) error {

	q := `SELECT deadline FROM todos WHERE id = $1 AND is_set = false`
	deadline := time.Time{}

	err := t.db.QueryRow(q, id).Scan(&deadline)
	if err != nil {
		return err
	}

	m := month.Month()
	fmt.Println("Month: ", int(m))
	deadline = deadline.AddDate(0, int(m), 0)
	q = `UPDATE todos SET deadline = $1 WHERE id = $2 AND is_set = false`
	_, err = t.db.Exec(q, deadline, id)
	if err != nil {
		return err
	}
	return nil
}

// Add set day to deadline of todo which is_set is false
// AddSetDayToDeadLine(id int64, day *time.Time) error
func (t *todoRepo) AddSetDayToDeadLine(id int64, day *time.Time) error {

	q := `SELECT deadline FROM todos WHERE id = $1 AND is_set = false`
	deadline := time.Time{}

	err := t.db.QueryRow(q, id).Scan(&deadline)
	if err != nil {
		return err
	}

	d := day.Day()
	fmt.Println("Day: ", d)
	deadline = deadline.AddDate(0, 0, d)
	q = `UPDATE todos SET deadline = $1 WHERE id = $2 AND is_set = false`
	_, err = t.db.Exec(q, deadline, id)
	if err != nil {
		return err
	}
	return nil
}

// Add set hour to deadline of todo which is_set is false
// AddSetHourToDeadLine(id int64, hour *time.Time) error
func (t *todoRepo) AddSetHourToDeadLine(id int64, hour *time.Time) error {

	q := `SELECT deadline FROM todos WHERE id = $1 AND is_set = false`
	deadline := time.Time{}

	err := t.db.QueryRow(q, id).Scan(&deadline)
	if err != nil {
		return err
	}

	h := hour.Hour()
	fmt.Println("Hour: ", h)
	deadline = deadline.Add(time.Duration(h) * time.Hour)
	q = `UPDATE todos SET deadline = $1 WHERE id = $2 AND is_set = false`
	_, err = t.db.Exec(q, deadline, id)
	if err != nil {
		return err
	}

	return nil
}

// Update is_set to true of todo
// SetIsSet(id int64, isSet bool) error
func (t *todoRepo) SetIsSet(id int64) error {
	q := `UPDATE todos SET is_set=true WHERE user_id = $1 AND is_set = false`
	_, err := t.db.Exec(q, id)
	if err != nil {
		return err
	}
	fmt.Println("Set is_set to true", id)
	return nil
}

/*
  Done(id string) error
  Delete(id string) error
*/

func (t *todoRepo) Done(id string) error {
	return nil
}

func (t *todoRepo) DeleteByID(id string) error {
	return nil
}

func (t *todoRepo) DeleteSetIsFalse(id int64) error {
	q := `DELETE FROM todos WHERE user_id = $1 AND is_set = false`
	_, err := t.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}

/*

	GetAllByUserID(id int64) ([]*model.Todo, error)
	GetNoSet(id int64) (*model.Todo, error)
	AddDescription(id string, description string) error
	AddDeadline(id string, deadline *time.Time) error
	AddPhotoURL(id string, photoURL string) error
	AddFileURL(id string, fileURL string) error

	AddSetYearToDeadLine(id string, year *time.Time) error
	AddSetMonthToDeadLine(id string, month *time.Time) error
	AddSetDayToDeadLine(id string, day *time.Time) error
	AddSetHourToDeadLine(id string, hour *time.Time) error

	SetIsSet(id string, isSet bool) error
	Done(id string) error
	Delete(id string) error

*/
