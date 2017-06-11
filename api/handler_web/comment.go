package web

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	db "github.com/mingrammer/meetup-api/api/database"

	"github.com/mingrammer/meetup-api/api/model"
	"github.com/mingrammer/meetup-api/api/serializer"
)

func GetAllComments(w rest.ResponseWriter, r *rest.Request) {
	eventID := r.PathParam("eid")
	event := GetEventOr404(eventID)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	comments := []model.Comment{}
	db.DBConn.Model(&event).Related(&comments)
	serializedComments := []serializer.CommentSerialzer{}
	for _, comment := range comments {
		serializedComments = append(serializedComments, *serializer.SerializeComment(&comment))
	}
	w.WriteJson(&serializedComments)
}

func CreateComment(w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	eventID := r.PathParam("eid")
	event := GetEventOr404(eventID)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	comment := model.Comment{}
	if err := r.DecodeJsonPayload(&comment); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := GetUserOr404(tokenString)
	comment.EventID = event.ID
	comment.WriterID = user.ID
	comment.WriterName = user.Name
	if err := db.DBConn.Save(&comment).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&comment)
}

func GetComment(w rest.ResponseWriter, r *rest.Request) {
	eventID := r.PathParam("eid")
	event := GetEventOr404(eventID)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	id := r.PathParam("cid")
	comment := GetCommentOr404(id)
	if comment == nil {
		rest.NotFound(w, r)
		return
	}
	serializedComment := serializer.SerializeComment(comment)
	w.WriteJson(&serializedComment)
}

func UpdateComment(w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	eventID := r.PathParam("eid")
	event := GetEventOr404(eventID)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	id := r.PathParam("cid")
	comment := GetCommentOr404(id)
	if comment == nil {
		rest.NotFound(w, r)
		return
	}
	user := GetUserOr404(tokenString)
	if comment.WriterID != user.ID {
		rest.Error(w, "This user is not writer of this comment", http.StatusForbidden)
		return
	}
	if err := r.DecodeJsonPayload(&comment); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.DBConn.Save(&comment).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&comment)
}

func DeleteComment(w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	eventID := r.PathParam("eid")
	event := GetEventOr404(eventID)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	id := r.PathParam("cid")
	comment := GetCommentOr404(id)
	if comment == nil {
		rest.NotFound(w, r)
		return
	}
	user := GetUserOr404(tokenString)
	if comment.WriterID != user.ID {
		rest.Error(w, "This user is not writer of this comment", http.StatusForbidden)
		return
	}
	if err := db.DBConn.Delete(&comment).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetCommentOr404 gets a comment instance if exists, or nil otherwise
func GetCommentOr404(id string) *model.Comment {
	comment := model.Comment{}
	if err := db.DBConn.First(&comment, id).Error; err != nil {
		return nil
	}
	return &comment
}
