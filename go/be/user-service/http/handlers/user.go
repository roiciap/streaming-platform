package handlers

import (
	"net/http"
	"regexp"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	handler := &UserHandler{}
	return handler
}

var (
	loginRegex    = regexp.MustCompile(`^\/login$`)
	registerRegex = regexp.MustCompile(`^\/register$`)
)

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && loginRegex.MatchString(r.URL.Path):
		h.login(w, r)
		return
	case r.Method == http.MethodPost && registerRegex.MatchString(r.URL.Path):
		h.register(w, r)
		return
	default:
		notFound(w, r)
	}
}

func (h *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
