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

func GetCategory(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	categoryID := r.PathParam("tid")
	category := GetCategoryOr404(db, categoryID)
	if category == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&category)
}

// GetCategoryOr404 gets a category instance if exists, or nil otherwise
func GetCategoryOr404(db *gorm.DB, id string) *model.Category {
	category := model.Category{}
	if err := db.First(&category, id).Error; err != nil {
		return nil
	}
	return &category
}
