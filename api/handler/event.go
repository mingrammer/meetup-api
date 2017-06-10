package handler

import (
	"net/http"

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
	userToken := r.Env["REMOTE_USER"]
	event := model.Event{}
	if err := r.DecodeJsonPayload(&event); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	event.OwnerToken = userToken.(string)
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
	id := r.PathParam("eid")
	event := GetEventOr404(db, id)
	if event == nil {
		rest.NotFound(w, r)
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
	id := r.PathParam("eid")
	event := GetEventOr404(db, id)
	if event == nil {
		rest.NotFound(w, r)
		return
	}
	if err := db.Delete(&event).Error; err != nil {
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
