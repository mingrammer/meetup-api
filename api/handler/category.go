package handler

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
)

func GetAllCategories(db *gorm.DB, w rest.ResponseWriter, _ *rest.Request) {
	categories := []model.Category{}
	db.Find(&categories)
	w.WriteJson(&categories)
}
