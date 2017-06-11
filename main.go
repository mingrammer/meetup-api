package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mingrammer/meetup-api/api"
	"github.com/mingrammer/meetup-api/config"
)

func main() {
	stop := make(chan bool)

	api.InitDB()
	webAPI := api.InitWebAPI()

	go func() {
		fmt.Printf("The meetup handler_web server is running on :%d\n", config.WebAPIConfig.Port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.WebAPIConfig.Port), webAPI.MakeHandler()))
	}()

	<-stop
}
