package templates

import (
	"net/http"
	"text/template"
)

// RenderTemplate renders the template along with data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
