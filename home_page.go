package main

import (
	"net/http"

	"github.com/pratikju/go-chat/session"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session.Manager.SessionDestroy(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}
