package postgres

import "github.com/jmoiron/sqlx"

type challengeRepo struct {
	db *sqlx.DB
}

func NewChallangeRepo(db *sqlx.DB) *challengeRepo {
	return &challengeRepo{
		db: db,
	}
}

func (c *challengeRepo) Create() {

}
func (c *challengeRepo) GetByID() {

}
func (c *challengeRepo) GetAll() {

}
func (c *challengeRepo) Update() {

}
func (c *challengeRepo) Delete() {

}
