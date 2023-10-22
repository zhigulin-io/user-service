package user

import (
	"database/sql"
	"errors"
	"log"
)

type PGStorage struct {
	db *sql.DB
}

func NewPGStorage() *PGStorage {
	db, err := sql.Open(
		"postgres",
		"user=zhigulin password=zhigulin dbname=userdb sslmode=disable",
	)

	if err != nil {
		log.Fatal("Cannot create database object:", err)
	}

	return &PGStorage{
		db: db,
	}
}

func (s *PGStorage) Write(username, password string) error {
	u := User{}
	row := s.db.QueryRow(FIND_BY_USERNAME_SQL, username)
	err := row.Scan(&u.Username, &u.Password)

	if err == nil {
		return errors.New("username already exists")
	}

	if err != sql.ErrNoRows {
		return err
	}

	_, err = s.db.Exec(INSERT_SQL, username, password)
	return err
}

const INSERT_SQL = `INSERT INTO users (username, password) VALUES (?, ?)`

const FIND_BY_USERNAME_SQL = `
	SELECT username, password 
	FROM users
	WHERE username = ?
`
