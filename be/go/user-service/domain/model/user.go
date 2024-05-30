package domain_model

import (
	"github.com/google/uuid"
)

type User struct {
	Id         int
	Username   string
	StreamGUID uuid.UUID
}
