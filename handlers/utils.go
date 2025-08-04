package handlers

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, "Not Found", 404)
}

func ErrorHandler(w http.ResponseWriter, text string, code int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"errorText": text,
		"status":    strconv.Itoa(code),
	}

	w.WriteHeader(code)
	tmpl.Execute(w, data)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("." + r.URL.Path)
	if err != nil {
		ErrorHandler(w, "Not found", 404)
	}

	w.Write(bytes)
}
