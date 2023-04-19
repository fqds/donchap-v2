package dto

type LobbyDto struct {
	ID              int    `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	MasterID        int    `json:"-" db:"master_id"`
	LobbyParameters []LobbyParameterDto
}
