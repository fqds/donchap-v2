package service

import (
	"java-to-go/config"
	"java-to-go/databaseConfig"
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/exception"
	"java-to-go/repository"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateUser(user *dto.UserDto) (string, error) {
	encryptedPassword, err := user.GetEncryptedPassword()
	if err != nil {
		return "", err
	}

	userToCreate := &entity.User{
		Name:              user.Name,
		EncryptedPassword: encryptedPassword,
	}
	id, err := repository.NewUserRep(databaseConfig.ConnectToDb()).CreateUser(userToCreate)
	if err != nil {
		return "", exception.NotCreatedObject{}
	}
	return id, nil
}

func CreateSession(user *dto.UserDto) (string, error) {

	userWithData := &entity.User{
		Name: user.Name,
	}

	if err := repository.NewUserRep(databaseConfig.ConnectToDb()).GetUserByName(userWithData); err != nil {
		return "", err
	}

	if !user.ComparePassword(userWithData.EncryptedPassword) {
		return "", exception.NotAuthenticated{}
	}

	expirationTime := time.Now().Add(time.Second * 900)
	claims := &Claims{
		UserID: userWithData.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Config.GetString("auth-secret")))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return signedToken, nil
}
