package ui

import (
	"embed"
	"greenlight/internal/data"
	"html/template"
	"log"
	"net/http"
)

//go:embed "templates"
var templateFS embed.FS

// var templates map[string]*template.Template

// do not put *.tpml or *.html files as templateNames
func HTML(templateName string) func(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Data data.User
	}
	// err :=  readJSON

	return func(w http.ResponseWriter, r *http.Request) {

		// read request
		tmpl, err := template.New(templateName).ParseFS(templateFS, "templates/"+templateName+".html")
		if err != nil {
			log.Println(err)
			// return err
		}

		err = tmpl.ExecuteTemplate(w, "home", input.Data)
		if err != nil {
			log.Println(err)
			// return err
		}
	}
}

// ExecuteTemplate basically
// i want to put the HTML function in cmd/api so the function readJSON will be accessible
// its precense only in the cmd/api prevents reading the Request body outside the main package
// its cool cuz it saves the 3-layer architecture
// so now we just kinda write the execute template function which will accept data to execute a template with.
