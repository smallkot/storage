package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/smallkot/storage/v2/internal/storage"
)

func main() {

	st := storage.NewStorage()
	file, err := st.Upload("text.txt", []byte("hello word"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It works", file)

	uuid, err1 := uuid.NewUUID()
	getFile, err := st.GetFile(uuid)
	if err != nil || err1 != nil {
		log.Fatal(err)
	}

	fmt.Println("It works", getFile)
}
