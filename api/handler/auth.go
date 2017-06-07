package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	"github.com/qkraudghgh/meetup/api/model"
	"github.com/qkraudghgh/meetup/api/oauth"
	"github.com/qkraudghgh/meetup/config"
)

func Authorize(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	config := config.GetConfig()

	if r.Method == "POST" {
		accessToken := r.PostForm.Get("access_token")
		user := GetUserOr404(db, accessToken)
		if user == nil {
			http.Redirect(w.(http.ResponseWriter), nil, config.SlackApp.RedirectURL, 302)
			return
		}
		http.Redirect(w.(http.ResponseWriter), nil, "/", 302)
		return
	}

	slackOauthCode := r.URL.Query().Get("code")
	slackTokenRequestURL := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
		oauth.SlackTokenURL,
		config.SlackApp.ClientId,
		config.SlackApp.ClientSecret,
		config.SlackApp.RedirectURL,
		slackOauthCode)

	resp, err := http.Get(slackTokenRequestURL)
	if err != nil {
		log.Print(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
	}

	slackOauthResp := oauth.SimpleSlackOauthResponse{}
	json.Unmarshal(body, &slackOauthResp)

	user := model.User{
		Token: slackOauthResp.AccessToken,
	}
	db.Save(&user)

	w.WriteJson(slackOauthResp)
}

// GetUserOr404 gets a user instance if exists, or nil otherwise
func GetUserOr404(db *gorm.DB, token string) *model.User {
	user := model.User{}
	if err := db.Where(&model.User{Token: token}).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
