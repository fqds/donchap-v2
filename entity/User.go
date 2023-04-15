package entity

type User struct {
	ID                string `json:"-" db:"id"`
	Name             string `json:"-" db:"name"`
	EncryptedPassword string `json:"-" db:"encrypted_password"`
}
