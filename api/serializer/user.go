package serializer

import (
	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/model"
)

type ParticipantSerialzer struct {
	ID        uint `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar"`
}

func SerializeParticipant(participant *model.User) *ParticipantSerialzer {
	owner := model.User{}
	db.DBConn.Find(&owner, participant.ID)
	participantSerialzer := ParticipantSerialzer{
		ID:        participant.ID,
		Name:      participant.Name,
		AvatarURL: participant.AvatarURL,
	}
	return &participantSerialzer
}
