package api

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/handler_web"
	"github.com/mingrammer/meetup-api/api/middleware"
)

type MeetupWebApp struct {
	API *rest.Api
}

func InitDB() {
	db.InitDB()
	db.InitSchema()
	db.InsertInitialData()
}

func InitWebAPI() *rest.Api {
	webAPI := rest.NewApi()
	webAPI.Use(rest.DefaultDevStack...)
	webAPI.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Origin",
			"Accept",
			"X-Requested-With",
			"Content-Type",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
			"Authorization",
		},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	webAPI.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			if request.URL.Path != "/auth" || request.URL.Path != "health-check" {
				return false
			}
			return true
		},
		IfTrue: &middleware.TokenAuthMiddleware{
			Realm: "meetup",
			Authenticator: func(token string) bool {
				user := web.GetUserOr404(token)
				if user != nil {
					return true
				}
				return false
			},
		},
	})
	router, err := rest.MakeRouter(
		rest.Get("/health-check", web.HealthCheck),
		rest.Get("/auth", web.Authorize),
		rest.Get("/events", web.GetAllEvents),
		rest.Post("/events", web.CreateEvent),
		rest.Get("/events/:eid", web.GetEvent),
		rest.Put("/events/:eid", web.UpdateEvent),
		rest.Delete("/events/:eid", web.DeleteEvent),
		rest.Get("/events/:eid/participants", web.GetAllParticipants),
		rest.Put("/events/:eid/join", web.JoinEvent),
		rest.Delete("/events/:eid/join", web.DisjoinEvent),
		rest.Get("/events/:eid/comments", web.GetAllComments),
		rest.Post("/events/:eid/comments", web.CreateComment),
		rest.Get("/events/:eid/comments/:cid", web.GetComment),
		rest.Put("/events/:eid/comments/:cid", web.UpdateComment),
		rest.Delete("/events/:eid/comments/:cid", web.DeleteComment),
		rest.Get("/categories", web.GetAllCategories),
		rest.Get("/categories/:tid", web.GetCategory),
	)
	if err != nil {
		log.Fatal(err)
	}
	webAPI.SetApp(router)
	return webAPI
}
