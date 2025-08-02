package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

// validate id & return it
func GetID(r *http.Request, path string) string {
	id := strings.TrimPrefix(r.URL.Path, "/"+path+"/")
	n, err := strconv.Atoi(id)

	if err != nil || id == "" || (n < 1 || n > 52) {
		return ""
	}

	return strconv.Itoa(n)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))).ServeHTTP(w, r)
}
