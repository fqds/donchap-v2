package entity

type User struct {
	ID       string `json:"-" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"-"`
}
