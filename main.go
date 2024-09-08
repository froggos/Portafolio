package main

import (
	"encoding/json"
	"log"
	"net/http"

	"portafolio/routes"
	"portafolio/structs"

	"github.com/gorilla/mux"
)

type ErrorAPI struct {
	Error     string
	ErrorCode string
	StateCode int
}

func logger(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		siguiente.ServeHTTP(w, r)
	})
}

func escribirJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func returnFuncHandleHTTP(t *structs.Templates) func(func(http.ResponseWriter, *http.Request, *structs.Templates) error) http.HandlerFunc {
	return func(f func(http.ResponseWriter, *http.Request, *structs.Templates) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if err := f(w, r, t); err != nil {
				escribirJSON(w, http.StatusBadRequest, &ErrorAPI{
					Error:     err.Error(),
					ErrorCode: "ERR_BAD_REQUEST",
					StateCode: http.StatusBadRequest,
				})
			}
		}
	}
}

func main() {
	mux := mux.NewRouter()

	mux.Use(logger)

	templateRender := structs.NewTemplate()

	if templateRender == nil {
		log.Fatalf("")
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
