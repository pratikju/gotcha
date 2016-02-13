package middleware

import (
	"net/http"

	"github.com/pratikju/go-chat/session"
)

// IsAuthenticated checks if user is authenticated for a specified route
func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := session.Manager.SessionStart(w, r)
		defer s.SessionRelease(w)

		if s.Get("profile") == nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		next(w, r)
	}
}
