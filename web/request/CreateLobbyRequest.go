package request

import "donchap-v2/dto"

type CreateLobbyRequest struct {
	Name            string                  `json:"name"`
	LobbyParameters []dto.LobbyParameterDto `json:"lobby_parameters"`
}
