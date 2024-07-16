package handlers

import (
	"html/template"
	"net/http"
)

func (c *ContactHandler) ContactForm(w http.ResponseWriter, r *http.Request) {
	tmpl := renderTemplate()
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderTemplate() *template.Template {
	tmpl, err := template.ParseFiles("ui/tmpl/index.html")
	if err != nil {
		panic(err)
	}
	return tmpl
}
