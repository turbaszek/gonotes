package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const line string = "=========="

func (env Env) parseFile(filePath string) {
	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}
	content := strings.TrimSpace(string(rawContent))

	rawNotes := strings.Split(content, line)

	for i := 0; i < len(rawNotes); i++ {
		err := env.parseAnCreateNote(rawNotes[i])
		if err != nil {
			log.Println(err)
		}
	}
}

func (env Env) parseAnCreateNote(rawNote string) error {
	parts := strings.Split(strings.TrimSpace(rawNote), "\n")
	if len(parts) != 4 {
		return fmt.Errorf("failed to parse a note")
	}
	bookName := strings.TrimSpace(parts[0])
	// _ := parts[1] information about note
	text := strings.TrimSpace(parts[3])
	book := env.addBook(bookName)
	env.addNote(text, book.ID)
	return nil
}
