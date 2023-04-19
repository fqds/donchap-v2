package service

import (
	"donchap-v2/databaseConfig"
	"donchap-v2/dto"
	"donchap-v2/entity"
	"donchap-v2/exception"
	"donchap-v2/repository"
)

func CreateLobby(lobby *dto.LobbyDto) error {
	lobbyToCreate := &entity.Lobby{
		Name:     lobby.Name,
		MasterID: lobby.MasterID,
	}
	
	if err := repository.NewLobbyRep(databaseConfig.ConnectToDb()).CreateLobby(lobbyToCreate); err != nil {
		return exception.NotCreatedObject{}
	}
	return nil
}