package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/projects/tgbot/config"
	"github.com/realtemirov/projects/tgbot/storage"
)

type Storege struct {
	db *sqlx.DB

	userRepo storage.UserI
	todoRepo storage.TodoI
}

func NewPostgres(cnf config.Config) (storage.StorageI, error) {
	psqlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.Postgres_HOST,
		cnf.Postgres_PORT,
		cnf.Postgres_USER,
		cnf.Postgres_PASS,
		cnf.Postgres_DBNAME,
		cnf.Postgres_SSLMODE,
	)

	db, err := sqlx.Open("postgres", psqlConnection)
	if err != nil {
		log.Fatalf("cannot connect to postgres: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("cannot connect to postgres: %s", err.Error())
	}

	return &Storege{
		db: db,
	}, nil
}

func (s *Storege) User() storage.UserI {
	if s.userRepo == nil {
		s.userRepo = NewUserRepo(s.db)
	}

	return s.userRepo
}

func (s *Storege) Todo() storage.TodoI {
	if s.todoRepo == nil {
		s.todoRepo = NewTodoRepo(s.db)
	}

	return s.todoRepo
}
