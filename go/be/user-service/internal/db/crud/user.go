package crud

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	db_model "github.com/roiciap/streaming-platform/go/be/user-service/internal/db/model"
)

const (
	insertUserQuery  = `INSERT INTO account.users (username, password_hash) VALUES ($1, $2) RETURNING id`
	readOneUserQuery = "SELECT id, username, password_hash,stream_guid, settings_id FROM users WHERE username='$1'"
)

func AddUser(user db_model.UserDbWrite) error {
	db, err := openDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var id int
	err = db.QueryRow(insertUserQuery, user.Username, user.PasswordHash).Scan(&id)
	if err != nil {
		log.Printf("Unable to insert data: %v\n", err)
		return err
	}

	return err
}

func ReadUserByLogin(login string) (*db_model.UserDbRead, error) {
	var user *db_model.UserDbRead
	var streamGUID sql.NullString
	db, err := openDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.QueryRow(readOneUserQuery, login).Scan(
		&user.Id, &user.Username, &user.PasswordHash, &streamGUID, &user.SettingsId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no user found with the given ID")
		}
		return nil, err
	}

	if streamGUID.Valid {
		user.StreamGUID, err = uuid.Parse(streamGUID.String)
		if err != nil {
			return nil, err
		}
	} else {
		user.StreamGUID = uuid.Nil
	}

	return user, err
}
