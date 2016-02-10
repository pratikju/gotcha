package main

import (
	"net/http"
	"text/template"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := GoChatManager.SessionStart(w, r)
	defer session.SessionRelease(w)

	profile := session.Get("profile")
	homeTemplate, err := template.New("webpage").Parse(homePage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = homeTemplate.Execute(w, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
