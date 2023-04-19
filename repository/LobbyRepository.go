package repository

import (
	"donchap-v2/entity"

	"github.com/jmoiron/sqlx"
)

type LobbyRepository interface {
	CreateLobby(lobby *entity.Lobby) error
}

type LobbyRep struct {
	LobbyRepository
}

func NewLobbyRep(db *sqlx.DB) *LobbyRep {
	return &LobbyRep{LobbyRepository: NewLobbyRepPostgres(db)}
}
