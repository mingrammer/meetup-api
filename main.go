package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mingrammer/meetup-api/api"
	"github.com/mingrammer/meetup-api/config"
)

func main() {
	config := config.GetConfig()
	api := api.Initialize(config)
	fmt.Printf("The meetup server is running on :%d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), api.MakeHandler()))
}
