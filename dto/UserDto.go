package dto

type UserDto struct {
	ID                string `json:"-" db:"id"`
	Login             string `json:"login" db:"login"`
	EncryptedPassword string `json:"-" db:"encrypted_password"`
}
