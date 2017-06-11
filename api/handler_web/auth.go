package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/model"
	"github.com/mingrammer/meetup-api/api/oauth"
	"github.com/mingrammer/meetup-api/config"
	"github.com/nlopes/slack"
)

func Authorize(w rest.ResponseWriter, r *rest.Request) {
	code := r.URL.Query().Get("code")
	redirectURI := r.URL.Query().Get("redirect_uri")
	if code == "" {
		w.WriteJson("Need a valid code")
		return
	}
	slackTokenURL := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		config.WebAPIConfig.SlackApp.TokenURL,
		config.WebAPIConfig.SlackApp.ClientID,
		config.WebAPIConfig.SlackApp.ClientSecret,
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
	token := slackOauthResp.AccessToken
	if token == "" {
		rest.Error(w, "There are invalid info for getting a token", http.StatusNonAuthoritativeInfo)
		return
	}
	user := GetUserOr404(token)
	if user == nil {
		userID, username, avatarURL := GetSlackUserProfileInfo(token)
		user = &model.User{
			Token:       token,
			SlackUserID: userID,
			Name:        username,
			AvatarURL:   avatarURL,
		}
		db.DBConn.Create(user)
	}
	w.WriteJson(user)
}

func GetSlackUserProfileInfo(token string) (string, string, string) {
	api := slack.New(token)
	userIdentity, _ := api.GetUserIdentity()
	userID := userIdentity.User.ID
	user, _ := api.GetUserInfo(userID)
	username := user.Name
	avatarURL := user.Profile.Image72
	return userID, username, avatarURL
}

// GetUserOr404 gets a user instance if exists, or nil otherwise
func GetUserOr404(token string) *model.User {
	user := model.User{}
	if err := db.DBConn.Where(&model.User{Token: token}).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
