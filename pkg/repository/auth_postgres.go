package repository

import (
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user courseGo.User) (int, error) {

	return 0, nil
}

func (r *AuthPostgres) Login(username, password string) (courseGo.User, error) {

	return courseGo.User{}, nil
}
