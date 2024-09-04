package routes

import (
	"log"
	"net/http"
	"text/template"
)

func DevolverInicio(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		log.Fatalf("No se pudo inicializar el template: %v", err)
		return err
	}

	tmpl.Execute(w, nil)

	return nil
}
