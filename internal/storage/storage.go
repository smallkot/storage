package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/smallkot/storage/v2/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetFile(fileId uuid.UUID) (*file.File, error) {
	file, ok := s.files[fileId]
	if !ok {
		return nil, fmt.Errorf("not found file. ID = %v", fileId)
	}
	return file, nil
}
