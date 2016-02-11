package main

import (
	"net/http"
	"text/template"

	"github.com/pratikju/go-chat/session"
	"github.com/pratikju/go-chat/templates"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := session.Manager.SessionStart(w, r)
	defer s.SessionRelease(w)

	profile := s.Get("profile")

	//TODO remove below code, use RenderTemplate from templates package
	homeTemplate, err := template.New("webpage").Parse(templates.HomePage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = homeTemplate.Execute(w, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
