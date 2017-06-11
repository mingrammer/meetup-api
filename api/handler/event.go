package handler

import (
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
	"github.com/mingrammer/meetup-api/api/serializer"
)

func GetAllEvents(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	dateStart := r.URL.Query().Get("start")
	dateEnd := r.URL.Query().Get("end")
	events := []model.Event{}
	selectedColumns := db.Select("id, title, date_start, date_end")
	switch {
	case dateStart != "" && dateEnd != "":
		selectedColumns.Where("date_start > ? AND date_end < ?", dateStart, dateEnd).Find(&events)
	case dateStart != "" && dateEnd == "":
		selectedColumns.Where("date_start > ?", dateStart).Find(&events)
	case dateStart == "" && dateEnd != "":
		selectedColumns.Where("date_end < ?", dateEnd).Find(&events)
	default:
		selectedColumns.Find(&events)
	}
	w.WriteJson(&events)
}

func CreateEvent(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	event := model.Event{}
	if err := r.DecodeJsonPayload(&event); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	category := GetCategoryOr404(db, strconv.FormatUint(uint64(event.CategoryID), 10))
	if category != nil {
		event.CategoryTitle = category.Title
	}
	user := GetUserOr404(db, tokenString)
	event.OwnerID = user.ID
	event.OwnerName = user.Name
	if err := db.Save(&event).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&event)
}

func GetEvent(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("eid")
	event := GetEventOr404(db, id)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	serializedEvent := serializer.SerializeEvent(db, event)
	w.WriteJson(&serializedEvent)
}

func UpdateEvent(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	id := r.PathParam("eid")
	event := GetEventOr404(db, id)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	user := GetUserOr404(db, tokenString)
	if event.OwnerID != user.ID {
		rest.Error(w, "This user is not owner of this event", http.StatusForbidden)
		return
	}
	if err := r.DecodeJsonPayload(&event); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.Save(&event).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&event)
}

func DeleteEvent(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	id := r.PathParam("eid")
	event := GetEventOr404(db, id)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	user := GetUserOr404(db, tokenString)
	if event.OwnerID != user.ID {
		rest.Error(w, "This user is not owner of this event", http.StatusForbidden)
		return
	}
	if err := db.Delete(&event).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllParticipants(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	users := []model.User{}
	db.Model(&event).Association("Participants").Find(&users)
	serializedParticipants := []serializer.ParticipantSerialzer{}
	for _, user := range users {
		serializedParticipants = append(serializedParticipants, *serializer.SerializeParticipant(db, &user))
	}
	w.WriteJson(&serializedParticipants)
}

func JoinEvent(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	user := GetUserOr404(db, tokenString)
	if err := db.Model(&event).Association("Participants").Append(user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DisjoinEvent(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	userToken := r.Env["VALID_USER_TOKEN"]
	tokenString := userToken.(string)
	eventId := r.PathParam("eid")
	event := GetEventOr404(db, eventId)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	user := GetUserOr404(db, tokenString)
	if err := db.Model(&event).Association("Participants").Delete(user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetEventOr404 gets a event instance if exists, or nil otherwise
func GetEventOr404(db *gorm.DB, id string) *model.Event {
	event := model.Event{}
	if err := db.First(&event, id).Error; err != nil {
		return nil
	}
	return &event
}
