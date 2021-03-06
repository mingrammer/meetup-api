package serializer

import (
	"time"

	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/model"
)

type WriterSerializer struct {
	ID        uint `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar"`
}

type CommentSerialzer struct {
	ID        uint `json:"id"`
	Content   string `json:"content"`
	EventID   uint `json:"event_id"`
	Writer    WriterSerializer `json:"writer"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func SerializeComment(comment *model.Comment) *CommentSerialzer {
	owner := model.User{}
	db.DBConn.Find(&owner, comment.WriterID)
	commentSerialzer := CommentSerialzer{
		ID:      comment.ID,
		Content: comment.Content,
		EventID: comment.EventID,
		Writer: WriterSerializer{
			ID:        owner.ID,
			Name:      owner.Name,
			AvatarURL: owner.AvatarURL,
		},
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
	return &commentSerialzer
}
