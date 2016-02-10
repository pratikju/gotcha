package main

import "net/http"

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	GoChatManager.SessionDestroy(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}
