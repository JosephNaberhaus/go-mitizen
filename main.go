package main

import (
	"github.com/JosephNaberhaus/go-mitizen/commit"
	"log"
	"os"
)

func main()  {
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	err = commit.Commit()
	if err != nil {
		log.Fatal(err)
	}
}