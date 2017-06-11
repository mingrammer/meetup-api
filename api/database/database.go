package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
	"github.com/mingrammer/meetup-api/config"
)

var DBConn *gorm.DB

func InitDB() {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DBConfig.Username,
		config.DBConfig.Password,
		config.DBConfig.Name,
		config.DBConfig.Charset)
	db, err := gorm.Open(config.DBConfig.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
		return
	}
	DBConn = db
}

func InitSchema() {
	DBConn.AutoMigrate(&model.User{})
	DBConn.AutoMigrate(&model.Category{})
	DBConn.AutoMigrate(&model.Event{})
	DBConn.Model(&model.Event{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	DBConn.Model(&model.Event{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	DBConn.AutoMigrate(&model.Comment{})
	DBConn.Model(&model.Comment{}).AddForeignKey("writer_id", "users(id)", "CASCADE", "CASCADE")
	DBConn.Model(&model.Comment{}).AddForeignKey("event_id", "events(id)", "CASCADE", "CASCADE")
}

func InsertInitialData() {
	categories := []model.Category{
		{Title: "drink"},
		{Title: "coffee"},
		{Title: "coding"},
		{Title: "talking"},
		{Title: "meal"},
	}
	for _, category := range categories {
		if err := DBConn.First(&model.Category{}, model.Category{Title: category.Title}).Error; err != nil {
			DBConn.Create(&category)
		}
	}
}
