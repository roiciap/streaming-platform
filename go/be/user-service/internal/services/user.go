package services

import (
	"encoding/json"
	"io"

	db_model "github.com/roiciap/streaming-platform/go/be/user-service/internal/db/model"
	http_model "github.com/roiciap/streaming-platform/go/be/user-service/internal/http/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

func ReadCredsFromRequest(body io.ReadCloser) (creds *http_model.CreditentialsRequest, err error) {
	err = json.NewDecoder(body).Decode(creds)
	if err != nil {
		return nil, err
	}
	if err := validator.Validate(*creds); err != nil {
		return nil, err
	}
	return nil, err
}

func BuildDbUserFromRequest(creds *http_model.CreditentialsRequest) (credsDb *db_model.UserDb, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 14)
	if err != nil {
		return nil, err
	}
	credsDb = &db_model.UserDb{
		Login:        creds.Login,
		PasswordHash: passwordHash,
	}
	return
}

func CheckPasswordMatch(password string, passwordHash []byte) bool {
	// maybe it should be more like this instead of bool:
	// if err == bcrypt.ErrMismatchedHashAndPassword...
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
