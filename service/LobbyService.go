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

	for _, value := range lobby.LobbyParameters {
		lobbyParameterToCreate := &entity.LobbyParameter{
			LobbyID:      lobby.ID,
			Name:         value.Name,
			Formula:      value.Formula,
			DropdownList: value.DropdownList,
			IsVisible:    value.IsVisible,
		}
		if err := repository.NewLobbyRep(databaseConfig.ConnectToDb()).CreateLobbyParameter(lobbyParameterToCreate); err != nil {
			return exception.NotCreatedObject{}
		}
	}

	return nil
}

func ConnectToLobby(lobbyName string, userID int) error {
	// lobby, err := repository.NewLobbyRep(databaseConfig.ConnectToDb()).GetLobbyByName(lobbyToCreate);
	// if err != nil {
	// 	return exception.NotCreatedObject{}
	// }

	return nil
}
