package db_model

type UserDb struct {
	Login        string
	PasswordHash []byte
}
