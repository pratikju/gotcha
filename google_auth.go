package main

import (
    "golang.org/x/oauth2"
    "net/http"
    "html/template"
    "io/ioutil"
)

var oauth_google_config = &oauth2.Config {
        ClientID: "112710392513-2mam2i72bj2lp045lge2fet5il10u48t.apps.googleusercontent.com",

        ClientSecret: "sVpHDvOwUf9am9Gvg5VQDa4i",

        Endpoint: oauth2.Endpoint{
    			AuthURL: "https://accounts.google.com/o/oauth2/auth",
    			TokenURL: "https://accounts.google.com/o/oauth2/token",
    		},

        RedirectURL: "http://localhost:8000/google_home",

        //This is the 'scope' of the data that you are asking the user's permission to access. For getting user's info, this is the url that Google has defined.
        Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
    }

//This is the URL that Google has defined so that an authenticated application may obtain the user's info in json format
const googleProfileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo"

func google_authorization_handler(w http.ResponseWriter, r *http.Request) {
    url := oauth_google_config.AuthCodeURL("")
    http.Redirect(w, r, url, http.StatusFound)
}

func google_home_handler(w http.ResponseWriter, r *http.Request){
    code := r.FormValue("code")
    token, err := oauth_google_config.Exchange(oauth2.NoContext, code)
    if err != nil {
      panic(err)
    }

    client := oauth_google_config.Client(oauth2.NoContext, token)
    response, _ := client.Get(googleProfileInfoURL)

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
