package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"portafolio/structs"
)

func GetProjects(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	err := templateRenderer.Render(w, "proyectos.html", nil)

	if err != nil {
		log.Println(err)
		return errors.New("ocurrio un error")
	}

	return nil
}

func GetRender(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	err := templateRenderer.Render(w, "render.html", nil)

	if err != nil {
		return fmt.Errorf("error al renderizar la plantilla: %w", err)
	}

	return nil
}
