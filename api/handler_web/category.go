package web

import (
	"github.com/ant0ine/go-json-rest/rest"
	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/model"
)

func GetAllCategories(w rest.ResponseWriter, _ *rest.Request) {
	categories := []model.Category{}
	db.DBConn.Find(&categories)
	w.WriteJson(&categories)
}

func GetCategory(w rest.ResponseWriter, r *rest.Request) {
	categoryID := r.PathParam("tid")
	category := GetCategoryOr404(categoryID)
	if category == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&category)
}

// GetCategoryOr404 gets a category instance if exists, or nil otherwise
func GetCategoryOr404(id string) *model.Category {
	category := model.Category{}
	if err := db.DBConn.First(&category, id).Error; err != nil {
		return nil
	}
	return &category
}
