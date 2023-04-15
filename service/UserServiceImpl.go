package service

import (
	"java-to-go/databaseConfig"
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/exception"
	"java-to-go/repository"
)

func CreateUser(user *entity.User) (string, error) {
	encryptedPassword, err := user.GetEncryptedPassword()
	if err != nil {
		return "", err
	}

	userToCreate := &dto.UserDto{
		Login:             user.Login,
		EncryptedPassword: encryptedPassword,
	}
	id, err := repository.NewUserRep(databaseConfig.ConnectToDb()).CreateUser(userToCreate)
	if err != nil {
		return "", exception.NotCreatedObject{}
	}
	return id, nil
}
