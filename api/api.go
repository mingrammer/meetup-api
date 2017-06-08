package api

import (
	"fmt"
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	"github.com/qkraudghgh/meetup/api/handler"
	"github.com/qkraudghgh/meetup/api/middleware"
	"github.com/qkraudghgh/meetup/api/model"
	"github.com/qkraudghgh/meetup/config"
)

type MeetupApp struct {
	Api    *rest.Api
	DB     *gorm.DB
	Config *config.Config
}

func Initialize(config *config.Config) *rest.Api {
	meetupApp := &MeetupApp{}
	meetupApp.Config = config
	meetupApp.InitDB()
	meetupApp.InitSchema()
	meetupApp.SetInitialData()

	tokenMiddleware := &middleware.TokenAuthMiddleware{
		Realm: "meetup",
		Authenticator: func(token string) bool {
			user := handler.GetUserOr404(meetupApp.DB, token)
			if user != nil {
				return true
			}
			return false
		},
	}

	meetupApp.Api = rest.NewApi()
	meetupApp.Api.Use(rest.DefaultDevStack...)
	meetupApp.Api.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != "/auth"
		},
		IfTrue: tokenMiddleware,
	})
	router, err := rest.MakeRouter(
		rest.Get("/auth", meetupApp.Authorize),
		rest.Get("/events", meetupApp.GetAllEvents),
		rest.Post("/events", meetupApp.CreateEvent),
		rest.Get("/events/:eid", meetupApp.GetEvent),
		rest.Put("/events/:eid", meetupApp.UpdateEvent),
		rest.Delete("/events/:eid", meetupApp.DeleteEvent),
		rest.Get("/events/:eid/comments", meetupApp.GetAllComments),
		rest.Post("/events/:eid/comments", meetupApp.CreateComment),
		rest.Get("/events/:eid/comments/:cid", meetupApp.GetComment),
		rest.Put("/events/:eid/comments/:cid", meetupApp.UpdateComment),
		rest.Delete("/events/:eid/comments/:cid", meetupApp.DeleteComment),
	)
	if err != nil {
		log.Fatal(err)
	}
	meetupApp.Api.SetApp(router)
	return meetupApp.Api
}

func (ma *MeetupApp) InitDB() {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		ma.Config.DB.Username,
		ma.Config.DB.Password,
		ma.Config.DB.Name,
		ma.Config.DB.Charset)
	db, err := gorm.Open(ma.Config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}
	ma.DB = db
}

func (ma *MeetupApp) InitSchema() {
	ma.DB.AutoMigrate(&model.User{})
	ma.DB.AutoMigrate(&model.Category{})
	ma.DB.AutoMigrate(&model.Event{})
	ma.DB.Model(&model.Event{}).AddForeignKey("owner_token", "users(token)", "CASCADE", "CASCADE")
	ma.DB.Model(&model.Event{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	ma.DB.AutoMigrate(&model.Comment{})
	ma.DB.Model(&model.Comment{}).AddForeignKey("writer_token", "users(token)", "CASCADE", "CASCADE")
	ma.DB.Model(&model.Comment{}).AddForeignKey("event_id", "events(id)", "CASCADE", "CASCADE")
}

func (ma *MeetupApp) SetInitialData() {
	categories := []model.Category{
		{Title: "drink"},
		{Title: "coffee"},
		{Title: "coding"},
		{Title: "talking"},
		{Title: "meal"},
	}
	for _, category := range categories {
		if err := ma.DB.First(&model.Category{}, model.Category{Title: category.Title}).Error; err != nil {
			ma.DB.Create(&category)
		}
	}
}

func (ma *MeetupApp) Authorize(w rest.ResponseWriter, r *rest.Request) {
	handler.Authorize(ma.DB, w, r)
}

func (ma *MeetupApp) GetAllEvents(w rest.ResponseWriter, r *rest.Request) {
	handler.GetAllEvents(ma.DB, w, r)
}

func (ma *MeetupApp) CreateEvent(w rest.ResponseWriter, r *rest.Request) {
	handler.CreateEvent(ma.DB, w, r)
}

func (ma *MeetupApp) GetEvent(w rest.ResponseWriter, r *rest.Request) {
	handler.GetEvent(ma.DB, w, r)
}

func (ma *MeetupApp) UpdateEvent(w rest.ResponseWriter, r *rest.Request) {
	handler.UpdateEvent(ma.DB, w, r)
}

func (ma *MeetupApp) DeleteEvent(w rest.ResponseWriter, r *rest.Request) {
	handler.DeleteEvent(ma.DB, w, r)
}

func (ma *MeetupApp) GetAllComments(w rest.ResponseWriter, r *rest.Request) {
	handler.GetAllComments(ma.DB, w, r)
}

func (ma *MeetupApp) CreateComment(w rest.ResponseWriter, r *rest.Request) {
	handler.CreateComment(ma.DB, w, r)
}

func (ma *MeetupApp) GetComment(w rest.ResponseWriter, r *rest.Request) {
	handler.GetComment(ma.DB, w, r)
}

func (ma *MeetupApp) UpdateComment(w rest.ResponseWriter, r *rest.Request) {
	handler.UpdateComment(ma.DB, w, r)
}

func (ma *MeetupApp) DeleteComment(w rest.ResponseWriter, r *rest.Request) {
	handler.DeleteComment(ma.DB, w, r)
}
