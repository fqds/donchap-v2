package entity

type User struct {
	ID                string `json:"-" db:"id"`
	Login             string `json:"-" db:"login"`
	EncryptedPassword string `json:"-" db:"encrypted_password"`
}
