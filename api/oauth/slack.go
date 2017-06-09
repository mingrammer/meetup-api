package oauth

type SimpleSlackOauthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}