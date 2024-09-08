package routes

import (
	"fmt"
	"net/http"
	"portafolio/structs"
)

func DevolverInicio(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {

	err := templateRenderer.Renderizar(w, "index.html", nil)

	if err != nil {
		return fmt.Errorf("error al renderizar la plantilla: %w", err)
	}

	return nil
}
