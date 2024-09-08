package structs

import (
	"io"
	"text/template"
)

type Templates struct {
	Template *template.Template
}

func (t *Templates) Render(w io.Writer, nombre string, datos interface{}) error {
	return t.Template.ExecuteTemplate(w, nombre, datos)
}

func NewTemplate() *Templates {
	return &Templates{
		Template: template.Must(template.ParseGlob("templates/*.html")),
	}
}
