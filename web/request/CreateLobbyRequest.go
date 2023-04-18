package request

import "donchap-v2/dto"

type CreateLobbyRequest struct {
	Name            string `json:"name"`
	MasterID        int    `json:"masterID"`
	LobbyParameters []dto.LobbyParameter
}
