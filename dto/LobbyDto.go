package dto

type LobbyDto struct {
	ID              int    `json:"-" db:"id"`
	Name            string `json:"-" db:"name"`
	MasterID        int    `json:"-" db:"master_id"`
	LobbyParameters []LobbyParameterDto
}
