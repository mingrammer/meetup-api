package handler

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
)

func GetAllComments(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	comments := []model.Comment{}
	db.Find(&comments)
	w.WriteJson(&comments)
}

func CreateComment(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["REMOTE_USER"]
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	event.OwnerToken = userToken.(string)
	comment := model.Comment{}
	if err := r.DecodeJsonPayload(&comment); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&comment)
}

func GetComment(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	id := r.PathParam("cid")
	comment := GetCommentOr404(db, id)
	if comment == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&event)
}

func UpdateComment(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	id := r.PathParam("cid")
	comment := GetCommentOr404(db, id)
	if comment == nil {
		rest.NotFound(w, r)
		return
	}
	if err := r.DecodeJsonPayload(&comment); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&comment)
}

func DeleteComment(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	id := r.PathParam("cid")
	comment := GetCommentOr404(db, id)
	if comment == nil {
		rest.NotFound(w, r)
		return
	}
	if err := db.Delete(&comment).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetCommentOr404 gets a comment instance if exists, or nil otherwise
func GetCommentOr404(db *gorm.DB, id string) *model.Comment {
	comment := model.Comment{}
	if err := db.First(&comment, id).Error; err != nil {
		return nil
	}
	return &comment
}
