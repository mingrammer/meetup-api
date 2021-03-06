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
	botAPI := api.InitBotAPI()

	go func() {
		fmt.Printf("The meetup web server is running on :%d\n", config.WebAPIConfig.Port)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", config.WebAPIConfig.Port), "certificate.crt", "private.key", webAPI.MakeHandler()))
	}()

	go func() {
		fmt.Printf("The meetup bot server is running on :%d\n", config.BotAPIConfig.Port)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", config.BotAPIConfig.Port), "certificate.crt", "private.key", botAPI.MakeHandler()))
	}()

	<-stop
}
