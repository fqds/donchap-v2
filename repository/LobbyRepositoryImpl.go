package repository

import (
	"donchap-v2/entity"

	"github.com/jmoiron/sqlx"
)

type LobbyRepPostgres struct {
	db *sqlx.DB
}

func (b LobbyRepPostgres) CreateLobby(lobby *entity.Lobby) error {
	if err := b.db.QueryRow("INSERT INTO lobbies (name, master_id) VALUES ($1, $2) RETURNING ID", lobby.Name, lobby.MasterID).Scan(&lobby.ID); err != nil {
		return err
	}

	return nil
}

func (b LobbyRepPostgres) CreateLobbyParameter(lobbyParameter *entity.LobbyParameter) error {
	if err := b.db.QueryRow("INSERT INTO lobby_parameters (lobby_id, name, formula, dropdown_list, is_visible) VALUES ($1, $2, $3, $4, $5) RETURNING ID", lobbyParameter.LobbyID, lobbyParameter.Name, lobbyParameter.Formula, lobbyParameter.DropdownList, lobbyParameter.IsVisible).Scan(&lobbyParameter.ID); err != nil {
		return err
	}

	return nil
}

func (b LobbyRepPostgres) GetLobbyByName(lobbyName string) (lobby entity.Lobby, err error) {
	lobby.Name = lobbyName
	err = b.db.QueryRow("SELECT id, master_id FROM lobbies WHERE name = $1", lobby.Name).Scan(&lobby.ID, &lobby.MasterID)
	return
}

func NewLobbyRepPostgres(db *sqlx.DB) *LobbyRepPostgres {
	return &LobbyRepPostgres{db: db}
}
