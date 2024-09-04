package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"portafolio/routes"

	"github.com/gorilla/mux"
)

type funcApi func(http.ResponseWriter, *http.Request) error

type ErrorAPI struct {
	Error        string
	CodigoError  string
	CodigoEstado int
}

type Template struct {
	template *template.Template
}

func logger(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		siguiente.ServeHTTP(w, r)
	})
}

func (t *Template) renderizar(w io.Writer, nombre string, datos interface{}) error {
	return t.template.ExecuteTemplate(w, nombre, datos)
}

// func nuevoTemplate() *Template {
// 	return &Template{
// 		template: template.Must(template.ParseGlob("*.html")),
// 	}
// }

func escribirJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func retornarFuncHandleHTTP(f funcApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			escribirJSON(w, http.StatusBadRequest, &ErrorAPI{
				Error:        err.Error(),
				CodigoError:  "ERR_BAD_REQUEST",
				CodigoEstado: http.StatusBadRequest,
			})
		}
	}
}

func main() {
	mux := mux.NewRouter()

	mux.Use(logger)

	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Rutas
	mux.HandleFunc("/", retornarFuncHandleHTTP(routes.DevolverInicio)).Methods("GET")
	mux.HandleFunc("/acerca-de", retornarFuncHandleHTTP(routes.DevolverAcercaDe)).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
