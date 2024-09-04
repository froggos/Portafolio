package routes

import (
	"log"
	"net/http"
	"text/template"
)

func DevolverAcercaDe(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("./templates/acerca.html")

	if err != nil {
		log.Fatalf("No se pudo inicializar el template> %v", err)
		return err
	}

	tmpl.Execute(w, nil)

	return nil
}
