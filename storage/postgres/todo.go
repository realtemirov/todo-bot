package postgres

import (
	"fmt"

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

func (t *todoRepo) Create(todo *model.Todo) (string, error) {
	fmt.Println(todo)
	var id string
	q := `INSERT INTO todos (id, created_at, updated_at,deleted_at, user_id, title, description, photo_url, file_url, deadline, is_done) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	err := t.db.QueryRow(
		q, todo.ID,
		todo.CreatedAt,
		todo.UpdatedAt,
		todo.DeletedAt,
		todo.User_ID,
		todo.Title,
		todo.Description,
		todo.Photo_URL,
		todo.File_URL,
		todo.Deadline,
		todo.IsDone).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil

}
func (t *todoRepo) GetByID(id string) (*model.Todo, error) {
	q := `SELECT user_id, title, description, photo_url, file_url FROM todos WHERE id = $1`
	todo := &model.Todo{}
	err := t.db.QueryRow(q, id).Scan(&todo.User_ID, &todo.Title, &todo.Description, &todo.Photo_URL, &todo.File_URL)

	if err != nil {
		return nil, err
	}
	return todo, nil

}
func (t *todoRepo) GetAllByUserID(id int64) ([]*model.Todo, error) {
	return nil, nil
}
func (t *todoRepo) Update(todo *model.Todo) (*model.Todo, error) {
	return nil, nil
}
func (t *todoRepo) Done(id string) error {
	return nil
}

func (t *todoRepo) Delete(id string) error {
	return nil
}
