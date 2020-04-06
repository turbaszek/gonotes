package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Env gives an easy access to database
type Env struct {
	gorm.DB
}

// Book is representation of a single book from kindle notes
type Book struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	BookHash string `gorm:"unique;not null"`
	Notes    []Note
}

// Note is representation of a single note from kindle notes
type Note struct {
	gorm.Model
	Text     string `gorm:"unique;not null"`
	NoteHash string
	BookID   uint
}

func (env *Env) addBook(name string) Book {
	var book Book
	bookHash := fmt.Sprintf("%x", sha256.Sum256([]byte(name)))
	book = Book{Name: name, BookHash: bookHash}
	env.DB.Unscoped().Where(book).Attrs(book).FirstOrCreate(&book)
	return book
}

func (env *Env) addNote(text string, bookID uint) uint {
	var note Note
	noteHash := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
	note = Note{Text: text, NoteHash: noteHash, BookID: bookID}
	env.DB.Unscoped().Where(note).Attrs(note).FirstOrCreate(&note)
	return note.ID
}

func (env *Env) updateBook(bookID uint, name string) {
	book := Book{}
	env.DB.Find(&book, bookID)
	book.Name = name
	env.DB.Save(&book)
}

func (env *Env) updateNote(noteID uint, text string) {
	note := Note{}
	env.DB.Find(&note, noteID)
	note.Text = text
	env.DB.Save(&note)
}

func (env *Env) removeBook(bookID uint) {
	var book Book
	env.First(&book, bookID)
	env.DB.Where("book_id = ?", bookID).Delete(Note{})
	env.DB.Delete(&book)
}

func (env *Env) removeNote(noteID uint) {
	env.DB.Delete(&Note{}, noteID)
}
