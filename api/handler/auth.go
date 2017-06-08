package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	"github.com/qkraudghgh/meetup/api/model"
	"github.com/qkraudghgh/meetup/api/oauth"
	"github.com/qkraudghgh/meetup/config"
)

func Authorize(db *gorm.DB, w rest.ResponseWriter, r *rest.Request) {
	config := config.GetConfig()

	code := r.URL.Query().Get("code")
	redirectURI := r.URL.Query().Get("redirect_uri")
	if code == "" {
		w.WriteJson("Need a valid code")
		return
	}
	slackTokenURL := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		config.SlackApp.TokenURL,
		config.SlackApp.ClientID,
		config.SlackApp.ClientSecret,
		code,
		redirectURI,
	)
	resp, err := http.Get(slackTokenURL)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slackOauthResp := oauth.SimpleSlackOauthResponse{}
	json.Unmarshal(body, &slackOauthResp)
	if slackOauthResp.AccessToken == "" {
		rest.Error(w, "There are invalid info for getting a token", http.StatusNonAuthoritativeInfo)
		return
	}
	user := GetUserOr404(db, slackOauthResp.AccessToken)
	if user == nil {
		db.Save(&model.User{
			Token: slackOauthResp.AccessToken,
		})
	}
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
