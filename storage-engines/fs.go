package storageengines

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gillesdemey/npm-registry/model"
)

type FSStorage struct {
	folder string
}

func NewFSStorage() *FSStorage {
	engine := new(FSStorage)
	engine.initialize()
	return engine
}

func (s *FSStorage) initialize() error {
	return nil
}

func (s *FSStorage) StoreTarball() error {
	return nil
}

func (s *FSStorage) RetrieveTarball(string, io.Writer) error {
	return nil
}

func (s *FSStorage) RetrieveUsernameFromToken(token string) (string, error) {
	tokenEntries := make(map[string]model.Token)
	if _, err := toml.DecodeFile("tokens.toml", &tokenEntries); err != nil {
		return "", err
	}
	tokenEntry := tokenEntries[token]

	return tokenEntry.Username, nil
}

func (s *FSStorage) StoreUser(pkg string) error {
	return nil
}

func (s *FSStorage) RetrieveUser(string, io.Writer) error {
	return nil
}

func (s *FSStorage) StoreUserToken(token string, username string) error {
	tokenEntry := make(map[string]model.Token, 1)
	tokenEntry[token] = model.Token{
		Username:  username,
		Timestamp: time.Now(),
	}

	entry := new(bytes.Buffer)
	if err := toml.NewEncoder(entry).Encode(tokenEntry); err != nil {
		return err
	}

	tokensFile := filepath.Join(s.folder, "tokens.toml")

	if err := ioutil.WriteFile(tokensFile, entry.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func (s *FSStorage) RetrieveMetadata(pkg string, writer io.Writer) error {
	metaFileName := fmt.Sprintf("packages/meta/%s.json", pkg)
	metaFileLocation := filepath.Join(s.folder, metaFileName)

	metaFile, err := os.Open(metaFileLocation)
	if err != nil {
		return err
	}
	defer metaFile.Close()
	io.Copy(writer, metaFile)

	return nil
}

func (s *FSStorage) StoreMetadata(pkg string, data io.Reader) error {
	metaFileName := fmt.Sprintf("packages/meta/%s.json", pkg)
	metaFileLocation := filepath.Join(s.folder, metaFileName)

	metaFile, err := os.Create(metaFileLocation)
	if err != nil {
		return err
	}
	defer metaFile.Close()
	io.Copy(metaFile, data)

	return nil
}