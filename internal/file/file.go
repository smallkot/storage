package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	uuid, er := uuid.NewUUID()
	if er != nil {
		return &File{}, er
	}

	return &File{
		ID:   uuid,
		Name: filename,
		Data: blob,
	}, nil
}

func (file *File) String() string {
	return fmt.Sprintf("File:\nID: %s\nName: %s\nData:%s", file.ID, file.Name, file.Data)
}
