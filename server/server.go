package server

import (
	"fmt"
	"log"
	"net/http"
)

// ListenHTTP starts http server at given hostname and port
func ListenHTTP(hostname string, port int, handler http.Handler) {
	host := fmt.Sprintf("%s:%d", hostname, port)
	log.Println("starting http server at", host)
	if err := http.ListenAndServe(host, handler); err != nil {
		log.Fatal(err)
	}
}

// AttachHandler attaches all http handler
func AttachHandler() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/logout", logoutHandler)
	// http.HandleFunc("/authorize_github", githubAuthorizationHandler)
	// http.HandleFunc("/git_home", gitHomeHandler)
	// http.HandleFunc("/authorize_google", googleAuthorizationHandler)
	// http.HandleFunc("/google_home", googleHomeHandler)
	// http.HandleFunc("/upload", uploadHandler)
	// http.HandleFunc("/uploads/", uploadViewHandler)
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	// http.Handle("/websocket", websocket.Handler(socketServer))
	// http.HandleFunc("/user", userHandler)
}