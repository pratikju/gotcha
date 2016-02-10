package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := oauthGoogleConfig.Client(oauth2.NoContext, token)
	response, err := client.Get(googleProfileInfoURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := GoChatManager.SessionStart(w, r)
	defer session.SessionRelease(w)

	session.Set("id_token", token.Extra("id_token"))
	session.Set("access_token", token.AccessToken)
	session.Set("profile", string(rawBody))

	http.Redirect(w, r, "/user", http.StatusMovedPermanently)

}
