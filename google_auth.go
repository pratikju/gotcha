package main

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "net/http"
    "html/template"
    "io/ioutil"
)

var oauth_google_config = &oauth2.Config {
        ClientID: "YOUR_GOOGLE_CLIENT_ID",

        ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",

        Endpoint: google.Endpoint,

        RedirectURL: "YOUR_GOOLE_REDIRECT_URL",

        Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
    }

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
