package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var oauthGitConfig = &oauth2.Config{
	ClientID: os.Getenv("GITHUB_CLIENT_ID"),

	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),

	RedirectURL: os.Getenv("GITHUB_REDIRECT_URL"),

	Endpoint: github.Endpoint,

	Scopes: []string{},
}

const githubProfileInfoURL = "https://api.github.com/user"

func githubAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthGitConfig.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func gitHomeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := oauthGitConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	client := oauthGitConfig.Client(oauth2.NoContext, token)
	response, _ := client.Get(githubProfileInfoURL)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	homeTemplate, err1 := template.New("webpage").Parse(homePage)
	if err1 != nil {
		panic(err1)
	}

	err2 := homeTemplate.Execute(w, string(body))
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}
