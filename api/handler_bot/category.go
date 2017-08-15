package bot

import (
	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/model"
)

func GetCategories() []model.Category {
	categories := []model.Category{}
	db.DBConn.Find(&categories)
	return categories
}
