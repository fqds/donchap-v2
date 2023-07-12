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
			LobbyID:      lobbyToCreate.ID,
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

func IsLobbyMaster(lobbyName string, userID int) (bool, error) {
	lobby, err := repository.NewLobbyRep(databaseConfig.ConnectToDb()).GetLobbyByName(lobbyName)
	if err != nil {
		err = exception.NotCreatedObject{}
		return false, err
	}

	if lobby.MasterID == userID {
		return true, nil
	}
	return false, nil
}

func ConnectToLobby(lobbyName string, userID int) (lobbyDto dto.LobbyDto, err error) {
	lobby, err := repository.NewLobbyRep(databaseConfig.ConnectToDb()).GetLobbyByName(lobbyName)
	if err != nil {
		err = exception.NotCreatedObject{}
		return
	}

	if lobby.MasterID == userID {
		lobbyDto, err = repository.NewLobbyRep(databaseConfig.ConnectToDb()).GetLobbyParameters(lobby.ID)
		return
	}
	
	// playerDto
	return
}
