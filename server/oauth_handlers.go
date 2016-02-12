package server

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/pratikju/go-chat/oauth/github"
	"github.com/pratikju/go-chat/oauth/google"
	"github.com/pratikju/go-chat/session"
)

//TODO check for CSRF Attack, pass the state to AuthCodeURL method
func googleAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	url := google.AuthConfig.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func githubAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	url := github.AuthConfig.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func githubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	config := github.AuthConfig
	profilesURL := github.ProfilesURL
	code := r.FormValue("code")
	handleCallback(w, r, config, profilesURL, code)
}

func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	config := google.AuthConfig
	profilesURL := google.ProfilesURL
	code := r.FormValue("code")
	handleCallback(w, r, config, profilesURL, code)
}

func handleCallback(w http.ResponseWriter, r *http.Request, config *oauth2.Config, profilesURL, code string) {
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := config.Client(oauth2.NoContext, token)
	response, _ := client.Get(profilesURL)

	defer response.Body.Close()
	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, _ := session.Manager.SessionStart(w, r)
	defer s.SessionRelease(w)

	s.Set("id_token", token.Extra("id_token"))
	s.Set("access_token", token.AccessToken)
	s.Set("profile", string(rawBody))

	http.Redirect(w, r, "/user", http.StatusMovedPermanently)
}
