package dto

type LobbyPlayerDto struct {
	ID       int  `json:"-" db:"id"`
	Position int  `json:"position" db:"position"`
	UserID   int  `json:"-" db:"user_id"`
	LobbyID  int  `json:"-" db:"lobby_id"`
	PlayerParameters []PlayerParameterDto
}
