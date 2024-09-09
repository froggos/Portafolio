package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"portafolio/routes"
	"portafolio/structs"

	"github.com/gorilla/mux"
)

type ErrorAPI struct {
	Error     string
	ErrorCode string
	StateCode int
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func returnFuncHandleHTTP(t *structs.Templates) func(func(http.ResponseWriter, *http.Request, *structs.Templates) error) http.HandlerFunc {
	return func(f func(http.ResponseWriter, *http.Request, *structs.Templates) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if err := f(w, r, t); err != nil {
				writeJSON(w, http.StatusBadRequest, &ErrorAPI{
					Error:     err.Error(),
					ErrorCode: "ERR_BAD_REQUEST",
					StateCode: http.StatusBadRequest,
				})
			}
		}
	}
}

func isSVG(fileName string) bool {
	ext := filepath.Ext(fileName)
	return strings.ToLower(ext) == ".svg"
}

func analyzeFile(directory string) error {
	files, err := os.ReadDir(directory)

	if err != nil {
		return err
	}

	for _, file := range files {
		if !isSVG(file.Name()) {
			if file.IsDir() {
				log.Printf("el archivo %s es una carpeta", file.Name())
				continue
			}

			log.Printf("advertencia: el archivo %s no es un SVG", file.Name())
		}
	}

	return nil
}

func main() {
	analyzeFile("./assets/img")

	mux := mux.NewRouter()

	mux.Use(logger)

	templateRender := structs.NewTemplate()

	if templateRender == nil {
		log.Fatalf("no existe render")
	}

	handleWithTemplates := returnFuncHandleHTTP(templateRender)

	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// rutas
	mux.HandleFunc("/", handleWithTemplates(routes.GetHome)).Methods("GET")
	mux.HandleFunc("/acerca-de", handleWithTemplates(routes.GetAbout)).Methods("GET")

	// metodos
	mux.HandleFunc("/acerca-de-habilidades", handleWithTemplates(routes.RedirectAboutAbilities)).Methods("GET")
	mux.HandleFunc("/acerca-de-mi", handleWithTemplates(routes.RedirectAboutMe)).Methods("GET")
	mux.HandleFunc("/acerca-de-experiencia", handleWithTemplates(routes.RedirectAboutExperience)).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
