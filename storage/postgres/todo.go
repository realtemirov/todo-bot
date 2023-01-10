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
	q := `INSERT INTO todos (id, created_at, user_id, text,  photo_url, file_url, deadline, is_set, is_done,notification)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err := t.db.QueryRow(q, todo.ID, todo.CreatedAt, todo.User_ID, todo.Text, todo.Photo_URL, todo.File_URL, todo.Deadline, todo.Is_Set, todo.Is_Done, todo.Notification).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Get todo by id  => user_id, title, description, photo_url, file_url
// GetByID(id string) (*model.Todo, error)
func (t *todoRepo) GetByID(id string) (*model.Todo, error) {

	q := `SELECT user_id, text,  photo_url, file_url, is_done, deadline, notification,is_set FROM todos WHERE id = $1 AND is_set = true`
	todo := model.Todo{}
	row := t.db.QueryRow(q, id)
	err := row.Scan(&todo.User_ID, &todo.Text, &todo.Photo_URL, &todo.File_URL, &todo.Is_Done, &todo.Deadline, &todo.Notification, &todo.Is_Set)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *todoRepo) GetNoSet(id int64) (*model.Todo, error) {
	var todo model.Todo

	q := `SELECT text, photo_url,file_url FROM todos WHERE is_set = false AND user_id = $1`
	row := t.db.QueryRow(q, id)
	err := row.Scan(&todo.Text, &todo.Photo_URL, &todo.File_URL)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// Get todo by user_id  => id, title, description, photo_url, file_url
// GetAllByUserID(id int64) ([]*model.Todo, error)
func (t *todoRepo) GetAllByUserID(id int64, done bool) ([]*model.Todo, error) {

	q := `SELECT id, text, photo_url, file_url, deadline,is_set,notification,is_done FROM todos WHERE user_id = $1 AND deleted_at IS NULL AND is_set = true AND is_done = $2 ORDER BY created_at DESC`
	todos := []*model.Todo{}
	err := t.db.Select(&todos, q, id, done)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepo) GetAllNotificationTimes() ([]*model.Notification, error) {
	q := `SELECT id, notification FROM todos WHERE deleted_at IS NULL AND is_set = true AND is_done = false AND EXTRACT(YEAR FROM deadline) =1`
	todos := []*model.Notification{}

	err := t.db.Select(&todos, q)

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepo) GetAllDeadlineTimes() ([]*model.Deadline, error) {
	q := `SELECT id, deadline FROM todos WHERE deleted_at IS NULL AND is_set = true AND is_done = false AND EXTRACT(YEAR FROM deadline) >1`
	todos := []*model.Deadline{}

	err := t.db.Select(&todos, q)

	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Add text to todos which is not set
// AddText(id int64, text string)
func (t *todoRepo) AddText(id int64, text string) error {
	var text1 string
	q := ` SELECT text FROM todos WHERE user_id=$1 AND is_set = false`

	row := t.db.QueryRow(q, id)
	err := row.Scan(&text1)
	if err != nil {
		return err
	}
	text1 += text + "!"
	q = `UPDATE todos SET text = $1 WHERE user_id = $2 AND is_set = false`
	_, err = t.db.Exec(q, text1, id)
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

// Add set date to deadline of todo which is_set is false
// AddTime(id int64, date *time.Time) error
func (t *todoRepo) AddTime(id int64, date *time.Time) error {

	var deadline time.Time
	q := `SELECT deadline FROM todos WHERE user_id = $1 AND is_set = false`

	err := t.db.QueryRow(q, id).Scan(&deadline)
	if err != nil {
		return err
	}

	if date.Year() != 1 {
		deadline = deadline.AddDate(date.Year()-1, int(date.Month())-1, date.Day()-1)
	} else {
		deadline = deadline.AddDate(0, 0, date.Day()-1)
	}

	q = `UPDATE todos SET deadline = $1 WHERE user_id = $2 AND is_set = false`
	_, err = t.db.Exec(q, deadline, id)
	if err != nil {
		return err
	}

	return nil
}

// Add set hour to deadline of todo which is_set is false
// AddSetHourToDeadLine(id int64, hour *time.Time) error
func (t *todoRepo) AddHour(id int64, hour *time.Duration, column string) (*model.Todo, error) {

	var (
		deadline time.Time
		q        string
	)
	if column == "deadline" {
		q = `SELECT deadline FROM todos WHERE user_id = $1 AND is_set = false`
	} else {
		q = `SELECT notification FROM todos WHERE user_id = $1 AND is_set = false`
	}

	err := t.db.QueryRow(q, id).Scan(&deadline)
	if err != nil {
		return nil, err
	}
	fmt.Println("deadline:", deadline)
	deadline = deadline.Add(*hour)
	if column == "deadline" {
		q = `UPDATE todos SET deadline = $1 WHERE user_id = $2 AND is_set = false `
	} else {
		q = `UPDATE todos SET notification = $1 WHERE user_id = $2 AND is_set = false`
	}
	q += ` RETURNING id, user_id, text, photo_url, file_url, deadline,is_set,notification,is_done`
	fmt.Println("deadline:", deadline)
	row := t.db.QueryRow(q, deadline, id)
	todo := &model.Todo{}
	err = row.Scan(&todo.ID, &todo.User_ID, &todo.Text, &todo.Photo_URL, &todo.File_URL, &todo.Deadline, &todo.Is_Set, &todo.Notification, &todo.Is_Done)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Update is_set to true of todo
// SetIsSet(id int64, isSet bool) error
func (t *todoRepo) SetIsSet(id int64) error {
	q := `UPDATE todos SET is_set=true WHERE user_id = $1 AND is_set = false`
	_, err := t.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}

/*
  Done(id string) error
  Delete(id string) error
*/

func (t *todoRepo) Done(id string, done bool) error {
	q := `UPDATE todos SET is_done = $1 WHERE id = $2`
	_, err := t.db.Exec(q, done, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepo) DeleteByID(id string) error {

	q := `UPDATE todos SET deleted_at = $1 WHERE id = $2`
	row, err := t.db.Exec(q, time.Now().Add(5*time.Hour), id)
	if err != nil {
		return err

	}

	if _, err = row.RowsAffected(); err != nil {
		fmt.Println("no")
	}
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
