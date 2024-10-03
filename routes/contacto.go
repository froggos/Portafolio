package routes

import (
	"errors"
	"log"
	"net/http"
	"net/smtp"
	"portafolio/structs"
)

func GetContact(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	err := templateRenderer.Render(w, "contacto.html", nil)

	if err != nil {
		log.Println(err)
		return errors.New("ocurrio un error")
	}

	return nil
}

func SendEmail(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
	r.ParseForm()

	clientName := r.FormValue("client-name")
	clientEmail := r.FormValue("client-email")
	clientMessage := r.FormValue("client-message")

	from := clientEmail
	to := "benjaminosoriovergara@gmail.com"
	subject := "Mensaje de contacto"
	body := "Nombre: " + clientName + "\nEmail: " + clientEmail + "\nMensaje:\n" + clientMessage

	msg := "Desde: " + from + "\n" + "Para: " + to + "\n" + "Asunto: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", to, "", "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))

	if err != nil {
		log.Println("entro al error")
		log.Println(err)

		w.Write([]byte("<p>Ocurrio un error al enviar el correo. Por favor intenta mas tarde.</p>"))

		return errors.New("ocurrio un error")
	}

	w.Write([]byte("<p>Gracias por tu mensaje, te responderé lo antes posible.</p>"))

	return nil
}
