package main

import (
    "golang.org/x/oauth2"
    "net/http"
    "html/template"
    "io/ioutil"
)

var oauth_git_config = &oauth2.Config {
        ClientID: "29aa81434e9e3fd3a196",

        ClientSecret: "73c6f3edf7b98a5b89875efa1fe2793a79179351",

        Endpoint: oauth2.Endpoint{
    			AuthURL: "https://github.com/login/oauth/authorize",
    			TokenURL: "https://github.com/login/oauth/access_token",
    		},

        RedirectURL: "http://localhost:8000/git_home",

        Scopes: []string{},
    }

//This is the URL that Google has defined so that an authenticated application may obtain the user's info in json format
const githubProfileInfoURL = "https://api.github.com/user"

func github_authorization_handler(w http.ResponseWriter, r *http.Request) {
    url := oauth_git_config.AuthCodeURL("")
    http.Redirect(w, r, url, http.StatusFound)
}

func git_home_handler(w http.ResponseWriter, r *http.Request){
    code := r.FormValue("code")
    token, err := oauth_git_config.Exchange(oauth2.NoContext, code)
    if err != nil {
      panic(err)
    }

    client := oauth_git_config.Client(oauth2.NoContext, token)
    response, _ := client.Get(githubProfileInfoURL)

    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)

    home_template, err1 := template.New("webpage").Parse(home_page)
    if err1 != nil {
      panic(err1)
    }

    err2 := home_template.Execute(w, string(body))
    if err2 != nil {
      http.Error(w, err2.Error(), http.StatusInternalServerError)
    }
}
