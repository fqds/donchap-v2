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
	if err := b.db.QueryRow("INSERT INTO users (name, encrypted_password) VALUES ($1, $2) RETURNING ID", user.Name, user.EncryptedPassword).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (b UserRepPostgres) GetUserByName(user *entity.User) error {
	if err := b.db.QueryRow("SELECT id, encrypted_password FROM users WHERE name = $1", user.Name).Scan(&user.ID, &user.EncryptedPassword); err != nil {
		return err
	}
	
	return nil
}

func NewUserRepPostgres(db *sqlx.DB) *UserRepPostgres {
	return &UserRepPostgres{db: db}
}
