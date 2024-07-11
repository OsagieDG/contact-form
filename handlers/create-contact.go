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
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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

	go func() {
		if err := sendEmail(name, email, phoneNumber, message); err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)

		}
	}()

	if _, err := w.Write([]byte("<h2>Thanks for reaching out!</h2>")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte("<h1>Message has been sent successfully!</h1>")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte("<h2>I will get back to you ASAP</h2>")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(`<p><strong>Click <a href="https://osag1e.dev/">osag1e.dev</a> to return to the homepage</strong></p>`)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
