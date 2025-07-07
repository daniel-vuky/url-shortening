package handlers

import (
	"net/http"

	"github.com/daniel-vuky/url-shortening/internal/services"
)

type Handler struct {
	service services.IURL
}

func NewHandler(service *services.URLService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RenderIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, I'm a shortener!"))
}
