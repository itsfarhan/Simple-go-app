// main.go
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Parse the templates once to avoid re-parsing on every request
var templates = template.Must(template.ParseGlob("templates/*.html"))

// RenderTemplate renders a specific template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Println("Template rendering error:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about")
}

func main() {
	http.HandleFunc("/", homeHandler)       // Route for home page
	http.HandleFunc("/about", aboutHandler) // Route for about page

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
