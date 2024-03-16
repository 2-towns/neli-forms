package page

import (
	"html/template"
	"log"
	"net/http"
)

func create(req *http.Request) *template.Template {
	tpl, err := template.New("").ParseGlob("./__resources__/static/templates/*.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	return tpl
}

// Display a template with data
func Display(w http.ResponseWriter, r *http.Request, path string, data interface{}) {
	tpl := create(r)
	err := tpl.ExecuteTemplate(w, path, data)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
