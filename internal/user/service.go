package user

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

type Service struct {
	storage UserStorage
}

type UserStorage interface {
	Write(username, password string) error
}

func NewService(storage UserStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) RegisterUser(username, password string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username cannot be empty")
	}

	if len(strings.TrimSpace(password)) < 8 {
		return errors.New("password cannot be less then 8 characters")
	}

	h := sha256.Sum256([]byte(password))
	password = hex.EncodeToString(h[:])

	return s.storage.Write(username, password)
}
