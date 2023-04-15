package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"-" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"-"`
}

func (u *User)GetEncryptedPassword() (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
