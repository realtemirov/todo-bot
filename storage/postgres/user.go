package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go/log"

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

	q := `INSERT INTO users (id, created_at) VALUES ($1, $2)`
	ok, err := r.db.Exec(q, user.ID, user.CreatedAt)

	if e(err) {
		return 0, err
	}

	affected, err := ok.RowsAffected()

	if e(err) {
		return 0, err
	}

	if affected == 0 {
		return 0, nil
	}

	return user.ID, nil
}

func (r *userRepo) Get(id int64) (*model.User, error) {

	user := model.User{}
	q := `SELECT id FROM users WHERE id=$1`
	row := r.db.QueryRow(q, id)
	err := row.Scan(&user.ID)

	if e(err) {
		return nil, err
	}

	err = row.Err()
	if e(err) {
		return nil, err
	}

	return &user, nil
}

func e(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	}
	return false
}
