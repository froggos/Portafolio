package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"portafolio/routes"
	"portafolio/structs"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
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

func returnFuncHandleHTTP(t *structs.Templates) func(func(http.ResponseWriter, *http.Request, *structs.Templates) error) http.HandlerFunc {
	return func(f func(http.ResponseWriter, *http.Request, *structs.Templates) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			if err := f(w, r, t); err != nil {
				http.Error(w, "Ocurrió un error, por favor intenta nuevamente.", http.StatusBadRequest)
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
	// err := structs.InitDb()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer structs.GetDb().Close()

	// structs.CreateMessagesTable()

	analyzeFile("./assets/img/logos")

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
	mux.HandleFunc("/contacto", handleWithTemplates(routes.GetContact)).Methods("GET")
	mux.HandleFunc("/proyectos", handleWithTemplates(routes.GetProjects)).Methods("GET")

	// metodos
	mux.HandleFunc("/acerca-de-habilidades", handleWithTemplates(routes.RedirectAboutAbilities)).Methods("GET")
	mux.HandleFunc("/acerca-de-mi", handleWithTemplates(routes.RedirectAboutMe)).Methods("GET")
	mux.HandleFunc("/acerca-de-experiencia", handleWithTemplates(routes.RedirectAboutExperience)).Methods("GET")
	// mux.HandleFunc("/enviar-email", handleWithTemplates(routes.SendEmail)).Methods("POST")
	mux.HandleFunc("/obtener-render", handleWithTemplates(routes.GetRender)).Methods("GET")

	log.Printf("servidor levantado en el puerto: %d", 80)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}
