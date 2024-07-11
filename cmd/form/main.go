package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/osag1e/contact-form/handlers"

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

	contactHandler := handlers.NewContactRequestHandler()

	router.Get("/", contactHandler.ContactForm)
	router.Post("/contact-form", contactHandler.HandleCreateContact)

	fmt.Println("Server is running on :4000")
	err := http.ListenAndServe(listenAddr, router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
