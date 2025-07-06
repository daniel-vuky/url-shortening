package handlers

import "net/http"

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, I'm a shortener!"))
}
