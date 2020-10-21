package http

import (
	"net/http"

	"github.com/cavke/go-chat-app"
)

type Handler struct {
	UserService chatapp.UserService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}
