package dto

import "golang.org/x/crypto/bcrypt"

type UserDto struct {
	ID       string `json:"-" db:"id"`
	Name    string `json:"-" db:"name"`
	Password string `json:"-" db:"password"`
}

func (u *UserDto) GetEncryptedPassword() (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
