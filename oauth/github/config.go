package github

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var AuthConfig = &oauth2.Config{
	ClientID: os.Getenv("GITHUB_CLIENT_ID"),

	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),

	RedirectURL: os.Getenv("GITHUB_REDIRECT_URL"),

	Endpoint: github.Endpoint,

	Scopes: []string{},
}

const ProfilesURL = "https://api.github.com/user"
