package routes

import (
	"fmt"
	"net/http"

	"portafolio/structs"
)

func GetAbout(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	tools := []structs.Tools{
		{
			Name:        "Go",
			ImageURL:    "./assets/img/logos/go.svg",
			Description: "Desarrollo de aplicaciones web, API Rest, websockets y SSE. Lenguaje de programación favorito.",
		},
		{
			Name:        "TypeScript",
			ImageURL:    "./assets/img/logos/typescript.svg",
			Description: "Segundo lenguaje de programación que más utilizo.",
		},
		{
			Name:        "Dart",
			ImageURL:    "./assets/img/logos/dart.svg",
			Description: "Desarrollo de aplicaciones móviles y de escritorio.",
		},
		{
			Name:        "PHP",
			ImageURL:    "./assets/img/logos/php.svg",
			Description: "Desarrollo de aplicaciones web monólito y API.",
		},
		{
			Name:        "Java",
			ImageURL:    "./assets/img/logos/java.svg",
			Description: "Desarrollo de aplicaciones de escritorio y web.",
		},
		{
			Name:        "JavaScript",
			ImageURL:    "./assets/img/logos/javascript.svg",
			Description: "Desarrollo de aplicaciones web utilizando distintos frameworks.",
		},
		{
			Name:        "C",
			ImageURL:    "./assets/img/logos/c.svg",
			Description: "Desarrollo de diversas aplicaciones, principalmente utilizado para modificar videojuegos.",
		},
		{
			Name:        "C#",
			ImageURL:    "./assets/img/logos/csharp.svg",
			Description: "Desarrollo aplicaciones de escritorio y web. Desarrollo de videojuegos con Unity.",
		},
		{
			Name:        "Angular",
			ImageURL:    "./assets/img/logos/angular.svg",
			Description: "Framework de JavaScript favorito y el que más utilizo.",
		},
		{
			Name:        "JQuery",
			ImageURL:    "./assets/img/logos/jquery.svg",
			Description: "Librería que utilicé bastante en un principio para el desarrollo frontend.",
		},
		{
			Name:        "Flutter",
			ImageURL:    "./assets/img/logos/flutter.svg",
			Description: "A la hora de desarrollar aplicaciones móviles es lo que más utilizo.",
		},
		{
			Name:        "Android Studio",
			ImageURL:    "./assets/img/logos/android.svg",
			Description: "Desarrollo de aplicaciones móviles.",
		},
		{
			Name:        "NodeJS",
			ImageURL:    "./assets/img/logos/node.svg",
			Description: "Desarrollo de aplicaciones diversas aplicaciones, principalmente web.",
		},
		{
			Name:        "MongoDB",
			ImageURL:    "./assets/img/logos/mongo.svg",
			Description: "Desarrollo de modelo de datos no relacional. Base de datos favorita.",
		},
		{
			Name:        "SQL Server",
			ImageURL:    "./assets/img/logos/sqlserver.svg",
			Description: "Gestor de bases de datos SQL que más domino.",
		},
		{
			Name:        "MySQL",
			ImageURL:    "./assets/img/logos/mysql.svg",
			Description: "Segundo gestor de bases de datos SQL que más domino.",
		},
		{
			Name:        "Spring Boot",
			ImageURL:    "./assets/img/logos/spring.svg",
			Description: "Herramienta que he usado bastante para el desarrollo de aplicaciones web de parte del backend.",
		},
		{
			Name:        "Ionic",
			ImageURL:    "./assets/img/logos/ionic.svg",
			Description: "Dominio moderado.",
		},
		{
			Name:        "Vue",
			ImageURL:    "./assets/img/logos/vue.svg",
			Description: "Segundo framework de JavaScript que más me gusta.",
		},
		{
			Name:        "Git",
			ImageURL:    "./assets/img/logos/git.svg",
			Description: "Herramienta para gestión de versiones de software que más utilizo.",
		},
		{
			Name:        "AWS",
			ImageURL:    "./assets/img/logos/aws.svg",
			Description: "Conocimiento moderado de los servicios que entrega AWS, principalmente utilizado para generar pipelines de despliegue.",
		},
		{
			Name:        "Google Cloud",
			ImageURL:    "./assets/img/logos/gcp.svg",
			Description: "Utilizado para levantar servidores diversos.",
		},
		{
			Name:        "Laravel",
			ImageURL:    "./assets/img/logos/laravel.svg",
			Description: "Framework que manejo pero no domino.",
		},
	}

	err := templateRenderer.Render(w, "acerca.html", tools)

	if err != nil {
		return fmt.Errorf("error al renderizar la plantilla: %w", err)
	}

	return nil
}
