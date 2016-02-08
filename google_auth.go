package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var oauthGoogleConfig = &oauth2.Config{
	ClientID: os.Getenv("GOOGLE_CLIENT_ID"),

	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),

	RedirectURL: os.Getenv("GOOGLE_REDIRECT_URL"),

	Endpoint: google.Endpoint,

	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
}

const googleProfileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo"

func googleAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthGoogleConfig.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func googleHomeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := oauthGoogleConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	client := oauthGoogleConfig.Client(oauth2.NoContext, token)
	response, _ := client.Get(googleProfileInfoURL)

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
