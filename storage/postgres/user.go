package postgres

import (
	"fmt"

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

	q := `INSERT INTO users (id, username, first_name, last_name,is_admin, photo_url, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	ok, err := r.db.Exec(q, user.ID, user.Username, user.FirsName, user.LastName, user.IsAdmin, user.Photo_URL, user.CreatedAt, user.UpdatedAt)

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

func (r *userRepo) Update(user *model.User) (*model.User, error) {

	q := `UPDATE users SET username=$1, first_name=$2, last_name=$3,  is_admin=$4, photo_url=$5, updated_at=$6 WHERE id=$7 RETURNING id, username, first_name, last_name, is_admin, photo_url, created_at, updated_at`
	row := r.db.QueryRow(q, user.Username, user.FirsName, user.LastName, user.IsAdmin, user.Photo_URL, user.UpdatedAt, user.ID)
	err := row.Scan(&user.ID, &user.Username, &user.FirsName, &user.LastName, &user.IsAdmin, &user.Photo_URL, &user.CreatedAt, &user.UpdatedAt)

	if e(err) {
		return nil, err
	}

	if err = row.Err(); e(err) {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user is null")
	}

	return user, nil
}

func (r *userRepo) Get(id int64) (*model.User, error) {

	var user *model.User
	q := `SELECT id, username, first_name, last_name, is_admin, photo_url, created_at, updated_at FROM users WHERE id=$1`
	row := r.db.QueryRow(q, id)
	err := row.Scan(&user.ID, &user.Username, &user.FirsName, &user.LastName, &user.IsAdmin, &user.Photo_URL, &user.CreatedAt, &user.UpdatedAt)

	if e(err) {
		return nil, err
	}

	err = row.Err()
	if e(err) {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetAll() ([]*model.User, error) {

	q := `SELECT id, username, first_name, last_name, is_admin, photo_url, created_at, updated_at FROM users`

	rows, err := r.db.Query(q)
	if e(err) {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.FirsName, &user.LastName, &user.IsAdmin, &user.Photo_URL, &user.CreatedAt, &user.UpdatedAt)
		if e(err) {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); e(err) {
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

func (r *userRepo) AddAdmin(id int64) error {
	q := `UPDATE users SET is_admin=$1 WHERE id=$2`
	ok, err := r.db.Exec(q, true, id)
	if e(err) {
		return err
	}

	i, err := ok.RowsAffected()
	if e(err) {
		return err
	}
	if i == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *userRepo) RemoveAdmin(id int64) error {
	q := `UPDATE users SET is_admin=$1 WHERE id=$2`
	ok, err := r.db.Exec(q, false, id)
	if e(err) {
		return err
	}
	i, err := ok.RowsAffected()
	if e(err) {
		return err
	}
	if i == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
