package routes

import (
	"errors"
	"log"
	"net/http"
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

// func SendEmail(w http.ResponseWriter, r *http.Request, templateRenderer *structs.Templates) error {
// 	r.ParseForm()

// 	log.Println("Test1")

// 	clientName := r.FormValue("client-name")
// 	clientEmail := r.FormValue("client-email")
// 	clientMessage := r.FormValue("client-message")

// 	message := structs.Message{
// 		Name:    clientName,
// 		Email:   clientEmail,
// 		Message: clientMessage,
// 	}

// 	err := message.InsertMessage()

// 	log.Println("Test2")

// 	if err != nil {
// 		log.Println("Entro al if del error")

// 		log.Println(err)

// 		w.Write([]byte("<p>Ocurrio un error al enviar el correo. Por favor intenta mas tarde.</p>"))

// 		return errors.New("ocurrio un error")
// 	}

// 	w.Write([]byte("<p>Gracias por tu mensaje, te responderé lo antes posible.</p>"))

// 	return nil
// }
