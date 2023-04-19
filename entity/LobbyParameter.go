package entity

type LobbyParameter struct {
	ID           int    `json:"-" db:"id"`
	LobbyID      int `json:"lobby_id" db:"lobby_id"`
	Name         string `json:"name" db:"name"`
	Formula      string `json:"formula" db:"formula"`
	DropdownList string `json:"dropdown_list" db:"dropdown_list"`
	IsVisible    bool   `json:"is_visible" db:"is_visible"`
}
