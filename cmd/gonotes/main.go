package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/turbaszek/gonotes/internal"
	"github.com/urfave/cli"
	"log"
	"os"
)

var home, _ = os.UserHomeDir()
var goNotesHome = fmt.Sprintf("%s/gonotes", home)

func main() {
	// Setup SQLite database
	err := os.Mkdir(goNotesHome, os.ModePerm)
	if os.IsNotExist(err) {
		panic(err)
	}
	db, err := gorm.Open("sqlite3", fmt.Sprintf("%s/notes.db", goNotesHome))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&internal.Book{}, &internal.Note{})
	env := &internal.Env{DB: db}

	// CLI app
	app := &cli.App{
		Name:                 "gonotes",
		Description:          "Simple tool to manage Kindle notes",
		EnableBashCompletion: true,
		Commands: []cli.Command{
			env.ParseNotesCmd(),
			{
				Name:      "book",
				Usage:     "Utilities to manage books",
				ShortName: "b",
				Subcommands: []cli.Command{
					env.ListBooksCmd(),
					env.DeleteBookCmd(),
					env.RemoveDuplicatesCmd(),
				},
			},
			env.ShowNotesCmd(),
			env.RandomNoteCmd(),
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
