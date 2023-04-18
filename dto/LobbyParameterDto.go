package dto

type LobbyParameter struct {
	ID           int    `json:"-" db:"id"`
	LobbyID      string `json:"-" db:"lobby_id"`
	Name         string `json:"name" db:"name"`
	Formula      string `json:"formula" db:"formula"`
	DropdownList string `json:"dropdown_list" db:"dropdown_list"`
	IsVisible    bool   `json:"is_visible" db:"is_visible"`
}
