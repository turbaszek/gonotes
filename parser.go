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

func (env Env) removeDuplicates(bookID uint) {
	var notes []Note
	env.DB.Where("book_id == ?", bookID).Find(&notes)

	i := 0
	for i < len(notes) {
		j := i + 1
		shift := true
		for j < len(notes) {
			if strings.Contains(notes[j].Text, notes[i].Text) {
				env.removeNote(notes[i].ID)
				notes = notes[1:]
				shift = false
				break
			} else {
				if strings.Contains(notes[i].Text, notes[j].Text) {
					env.removeNote(notes[j].ID)
					notes[j] = notes[len(notes)-1]
				} else {
					j++
				}
			}
		}
		if shift {
			i++
		}
	}
}
