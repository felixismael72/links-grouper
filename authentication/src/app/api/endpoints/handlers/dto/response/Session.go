package response

type Session struct {
	AccessToken string `json:"access_token"`
}

func NewSession(accessToken string) *Session {
	return &Session{AccessToken: accessToken}
}
