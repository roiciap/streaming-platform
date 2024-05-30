package handlers

import (
	"net/http"
	"regexp"

	"github.com/roiciap/streaming-platform/be/go/user-service/internal/db/crud"
	"github.com/roiciap/streaming-platform/be/go/user-service/internal/services"
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
	// read user from db and compare password
	user, err := crud.ReadUserByLogin(creds.Login)
	if err != nil {
		http.Error(w, "Couldnt find user", http.StatusBadRequest)
		return
	}
	match := services.CheckPasswordMatch(creds.Password, user.PasswordHash)
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
	// build crud-add user model
	user, err := services.BuildUserWriteFromRequest(creds)
	if err != nil {
		http.Error(w, "Problem occured in creating user", http.StatusInternalServerError)
		return
	}
	// save user to DB
	err = crud.AddUser(*user)
	if err != nil {
		http.Error(w, "Couldnt save user", http.StatusInternalServerError)
		return
	}
	// add cookies

	w.WriteHeader(http.StatusOK)
}
