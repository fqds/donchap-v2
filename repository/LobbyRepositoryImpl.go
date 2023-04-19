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

func NewLobbyRepPostgres(db *sqlx.DB) *LobbyRepPostgres {
	return &LobbyRepPostgres{db: db}
}
