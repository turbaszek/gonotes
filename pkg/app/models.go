package app

import (
	"github.com/jinzhu/gorm"
)

// Env gives an easy access to database
type Env struct {
	Db *gorm.DB
}

// Book is representation of a single book from kindle notes
type Book struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Notes []Note
}

// Note is representation of a single note from kindle notes
type Note struct {
	gorm.Model
	Text   string `gorm:"unique;not null"`
	BookID uint
}
