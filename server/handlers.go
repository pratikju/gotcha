package server

import (
	"net/http"
	"text/template"

	"github.com/pratikju/go-chat/templates"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginTemplate, err := template.New("webpage").Parse(templates.LoginPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = loginTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
