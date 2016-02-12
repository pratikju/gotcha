package templates

import (
	"net/http"
	"text/template"
)

// Render renders the template along with data
func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
