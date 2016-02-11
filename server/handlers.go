package server

import (
	"net/http"

	"github.com/pratikju/go-chat/templates"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, templates.LoginPage, nil)
}
