package handlers

import (
	"log"
	"net/http"
	"os"
)

func (c *ContactHandler) ContactForm(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "ui/tmpl/index.html")
}

func serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.Open(filePath)
	checkError(w, err)
	defer file.Close()

	fi, err := file.Stat()
	checkError(w, err)

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)
}

func checkError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error opening file:", err)
		return
	}
}
