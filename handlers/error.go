package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

// NotFound handles 404 errors by displaying a custom "page not found" message.
func NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, "This page does not exist.", 404)
}

// ErrorHandler renders a custom error page with the provided text and status code.
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
