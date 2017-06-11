package serializer

import (
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
)

type ParticipantSerialzer struct {
	ID        uint `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar"`
}

func SerializeParticipant(db *gorm.DB, participant *model.User) *ParticipantSerialzer {
	owner := model.User{}
	db.Find(&owner, participant.ID)
	participantSerialzer := ParticipantSerialzer{
		ID:        participant.ID,
		Name:      participant.Name,
		AvatarURL: participant.AvatarURL,
	}
	return &participantSerialzer
}
