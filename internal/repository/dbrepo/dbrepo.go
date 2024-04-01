package dbrepo

import (
	"database/sql"

	"github.com/onuraltuntasb/e-commerce/internal/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgressRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}
