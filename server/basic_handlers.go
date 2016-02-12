package server

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/pratikju/go-chat/session"
	"github.com/pratikju/go-chat/templates"
)

// File contains info about the type of file being uploaded
type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Files is a list of File
type Files []File

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, templates.LoginPage, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session.Manager.SessionDestroy(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := session.Manager.SessionStart(w, r)
	defer s.SessionRelease(w)

	profile := s.Get("profile")
	templates.Render(w, templates.HomePage, profile)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("files")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	files := Files{
		File{Name: handler.Filename, Type: handler.Header["Content-Type"][0]},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(files); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func uploadViewHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
