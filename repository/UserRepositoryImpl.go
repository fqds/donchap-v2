package repository

import (
	"java-to-go/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepPostgres struct {
	db *sqlx.DB
}

func (b UserRepPostgres) CreateUser(user *entity.User) (string, error) {
	var id string
	if err := b.db.QueryRow("INSERT INTO users (login, encrypted_password) VALUES ($1, $2) RETURNING ID", user.Login, user.EncryptedPassword).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func NewUserRepPostgres(db *sqlx.DB) *UserRepPostgres {
	return &UserRepPostgres{db: db}
}
