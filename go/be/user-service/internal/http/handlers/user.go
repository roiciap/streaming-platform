package handlers

import (
	"net/http"
	"regexp"

	"github.com/roiciap/streaming-platform/go/be/user-service/internal/services"
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
	creds, err := services.ReadCredsFromRequest(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	match := services.CheckPasswordMatch(creds.Password, []byte("AAA"))
	if !match {
		http.Error(w, "Invalid creditentials", http.StatusBadRequest)
		return
	}

	// add cookies

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	creds, err := services.ReadCredsFromRequest(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = services.BuildDbUserFromRequest(creds)
	if err != nil {
		http.Error(w, "Problem occured in creating user", http.StatusInternalServerError)
		return
	}
	// save creds to DB

	// add cookies
	w.WriteHeader(http.StatusOK)
}
