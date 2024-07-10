package handlers

import (
	"net/http"
)

type ContactHandler struct {
}

func NewContactRequestHandler() *ContactHandler {
	return &ContactHandler{}
}

func (c *ContactHandler) HandleCreateContact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phoneNumber := r.FormValue("phonenumber")
	message := r.FormValue("message")

	if name == "" || email == "" || phoneNumber == "" || message == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if err := sendEmail(name, email, phoneNumber, message); err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("<h2>Message has been sent successfully!</h2>"))
}
