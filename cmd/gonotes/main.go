package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/turbaszek/gonotes/internal"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const version = "v0.0.1"
const appName = "gonotes"

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
		Name:                 appName,
		Version:              version,
		Usage:                "Simple tool to manage Kindle notes",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			env.NewParseNotesCmd(),
			{
				Name:    "book",
				Usage:   "Utilities to manage books",
				Aliases: []string{"b"},
				Subcommands: []*cli.Command{
					env.NewListBooksCmd(),
					env.NewDeleteBookCmd(),
					env.NewRemoveDuplicatesCmd(),
				},
			},
			env.NewShowNotesCmd(),
			env.NewRandomNoteCmd(),
			internal.NewCompleteCommand(),
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
