package repository

import (
	"java-to-go/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *entity.User) (string, error)
}

type UserRep struct {
	UserRepository
}

func NewUserRep(db *sqlx.DB) *UserRep {
	return &UserRep{UserRepository: NewUserRepPostgres(db)}
}
