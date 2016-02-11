package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pratikju/go-chat/session"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := oauthGitConfig.Client(oauth2.NoContext, token)
	response, _ := client.Get(githubProfileInfoURL)

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
