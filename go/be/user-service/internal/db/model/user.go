package db_model

import "github.com/google/uuid"

type UserDbWrite struct {
	Username     string
	PasswordHash []byte
}

type UserDbRead struct {
	Id           int
	Username     string
	PasswordHash []byte
	StreamGUID   uuid.UUID
	SettingsId   int
}
