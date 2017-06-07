package oauth

const (
	SlackAuthURL  = "https://slack.com/oauth/authorize"
	SlackTokenURL = "https://slack.com/api/oauth.access"
)

type SimpleSlackOauthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}
