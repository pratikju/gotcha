package google

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var AuthConfig = &oauth2.Config{
	ClientID: os.Getenv("GOOGLE_CLIENT_ID"),

	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),

	RedirectURL: os.Getenv("GOOGLE_REDIRECT_URL"),

	Endpoint: google.Endpoint,

	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
}

const ProfilesURL = "https://www.googleapis.com/oauth2/v1/userinfo"
