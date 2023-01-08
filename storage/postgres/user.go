package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go/log"
	"github.com/spf13/cast"

	"github.com/realtemirov/projects/tgbot/model"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Add(user *model.User) (int64, error) {
	q := `INSERT INTO users ( id , action, created_at ) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(q, cast.ToString(user.ID), user.Action, user.CreatedAt)

	if e(err) {
		return 0, err
	}

	return user.ID, nil
}

func (r *userRepo) Get(id int64) (*model.User, error) {

	user := model.User{}
	q := `SELECT id, created_at FROM users WHERE id=$1`
	row := r.db.QueryRow(q, id)
	err := row.Err()
	if e(err) {
		return nil, err
	}

	err = row.Scan(&user.ID, &user.CreatedAt)
	if e(err) {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) SetAction(id int64, action string) error {

	q := `UPDATE users SET action = $1 WHERE id = $2`
	_, err := r.db.Exec(q, action, id)
	if e(err) {
		return err
	}

	return nil
}

func (r *userRepo) GetAction(id int64) (string, error) {

	var action string
	q := `SELECT action FROM users WHERE id = $1`
	err := r.db.QueryRow(q, id).Scan(&action)
	if e(err) {
		return "", err
	}

	return action, nil
}

func (r *userRepo) GetAll() ([]*model.User, error) {

	q := `SELECT id, action, created_at FROM users`
	users := []*model.User{}
	err := r.db.Select(&users, q)
	if e(err) {
		return nil, err
	}

	return users, nil
}

func e(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	}
	return false
}
