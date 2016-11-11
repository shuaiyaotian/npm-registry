package storageengines

import (
	"github.com/gillesdemey/npm-registry/model"
)

type GoogleStorage struct {
	bucket string
}

func NewGoogleStorage() *GoogleStorage {
	engine := new(GoogleStorage)
	engine.initialize()
	return engine
}

// 1. check if bucket exists
// 2. create bucket if it does not exist
func (s *GoogleStorage) initialize() error {
	return nil
}

func (s *GoogleStorage) StoreTarball() error {
	return nil
}

func (s *GoogleStorage) RetrieveTarball() ([]byte, error) {
	return nil, nil
}

func (s *GoogleStorage) RetrieveUser(token string) (model.User, error) {
	return model.User{}, nil
}

func (s *GoogleStorage) StoreUserToken(token string, username string) error {
	return nil
}