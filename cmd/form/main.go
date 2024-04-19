package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")

	router.Get("/", indexHandler)

	router.Post("/contact-form", contactHandler)

	fmt.Println("Server is running on :4000")
	err := http.ListenAndServe(listenAddr, router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/tmpl/index.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	phoneNumber := r.FormValue("phonenumber")
	message := r.FormValue("message")

	// Do something with the form data (e.g., send an email, save to database, etc.)
	fmt.Printf("Received message from: %s <%s>\n", name, email)
	fmt.Printf("Phone Number: %s\n", phoneNumber)
	fmt.Printf("Message: %s\n", message)

	// Send a response back to the client
	fmt.Fprintf(w, "Message received successfully!")
}
