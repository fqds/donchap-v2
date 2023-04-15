package repository

import (
	"java-to-go/dto"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *dto.UserDto) (string, error)
}

type UserRep struct {
	UserRepository
}

func NewUserRep(db *sqlx.DB) *UserRep {
	return &UserRep{UserRepository: NewUserRepPostgres(db)}
}
