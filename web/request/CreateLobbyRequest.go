package request

import "java-to-go/dto"

type CreateLobbyRequest struct {
	Name            string `json:"name"`
	MasterID        int    `json:"masterID"`
	LobbyParameters []dto.LobbyParameter
}
