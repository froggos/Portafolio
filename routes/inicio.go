package routes

import (
	"fmt"
	"net/http"
	"portafolio/structs"
)

func GetHome(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {

	err := templateRenderer.Render(w, "index.html", nil)

	if err != nil {
		return fmt.Errorf("error al renderizar la plantilla: %w", err)
	}

	return nil
}

func RedirectAboutAbilities(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	http.Redirect(w, r, "/acerca-de#habilidades-container", http.StatusSeeOther)

	return nil
}

func RedirectAboutMe(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	http.Redirect(w, r, "/acerca-de#yo-container", http.StatusSeeOther)

	return nil
}

func RedirectAboutExperience(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	http.Redirect(w, r, "/acerca-de#experiencia-container", http.StatusSeeOther)

	return nil
}
