package domain_model

import (
	"github.com/google/uuid"
)

type User struct {
	Id         int
	Login      string
	StreamGuid uuid.UUID
}
