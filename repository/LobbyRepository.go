package repository

import (
	"donchap-v2/dto"
	"donchap-v2/entity"

	"github.com/jmoiron/sqlx"
)

type LobbyRepository interface {
	CreateLobby(lobby *entity.Lobby) error
	CreateLobbyParameter(lobby *entity.LobbyParameter) error
	GetLobbyByName(lobbyName string) (entity.Lobby, error)
	GetLobbyParameters(lobbyID int) (dto.LobbyDto, error)
}

type LobbyRep struct {
	LobbyRepository
}

func NewLobbyRep(db *sqlx.DB) *LobbyRep {
	return &LobbyRep{LobbyRepository: NewLobbyRepPostgres(db)}
}
