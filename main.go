package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

var home, _ = os.UserHomeDir()
var gonotesHome = fmt.Sprintf("%s/gonotes", home)

const ls = "ls"
const rm = "rm"

func strToUint(n string) (uint, error) {
	v, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func (env *Env) listBooks() cli.Command {
	return cli.Command{
		Name:        ls,
		ArgsUsage:   "",
		Description: "Lists all books",
		Usage:       "Lists books",
		Action: func(c *cli.Context) {
			var books []Book
			env.DB.Find(&books)
			if len(books) == 0 {
				fmt.Println("It seems you have no books yet. Try to parse clippings from your Kindle.")
				return
			}

			templates := &promptui.SelectTemplates{
				Inactive: "{{ .ID }} | {{ .Name }}",
				Active:   "{{ .ID | cyan }} | {{ .Name | cyan }}",
				Selected: "{{ .Name | bold }}",
			}

			prompt := promptui.Select{
				Label:     "Your books",
				Items:     books,
				Templates: templates,
				Size:      8,
			}

			idx, _, err := prompt.Run()
			if err != nil {
				return
			}
			env.showNotes(books[idx].ID, true)
		},
	}
}

func (env *Env) deleteBook() cli.Command {
	return cli.Command{
		Name:        rm,
		Description: "Deletes book by ID",
		ArgsUsage:   "BOOK_ID",
		Usage:       "Deletes a book",
		Action: func(c *cli.Context) {
			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return
			}
			env.removeBook(bookID)
		},
	}
}

func (env *Env) showNotes(bookID uint, index bool) {
	var notes []Note
	env.DB.Where("book_id = ?", bookID).Find(&notes)
	if len(notes) == 0 {
		fmt.Println("No notes :<")
		return
	}
	for i := 0; i < len(notes); i++ {
		n := notes[i]
		if index {
			fmt.Println(fmt.Sprintf("%d) %s\n", n.ID, n.Text))
		} else {
			fmt.Println(fmt.Sprintf("%s\n", n.Text))
		}
	}
}

func (env *Env) listNotes() cli.Command {
	var useIndex bool
	return cli.Command{
		Name:        "notes",
		Description: "Lists all notes from provided book",
		ArgsUsage:   "BOOK_ID",
		Usage:       "List notes",
		ShortName:   "n",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "index, id", Destination: &useIndex, Usage: "If set notes id will be included"},
		},
		Action: func(c *cli.Context) {
			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return
			}
			env.showNotes(bookID, useIndex)
		},
	}
}

func (env *Env) randomNoteCmd() cli.Command {
	var asQuote bool
	var lenLimit int
	return cli.Command{
		Name:        "random",
		Description: "Shows a random note",
		Usage:       "Shows a random note",
		ShortName:   "r",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "quote, q", Destination: &asQuote, Usage: "Include book title and author"},
			cli.IntFlag{Name: "length, l", Destination: &lenLimit, Value: -1, Usage: "Limit note length to this value"},
		},
		Action: func(c *cli.Context) {
			note := env.getRandomNote(lenLimit)
			if asQuote {
				book := Book{}
				env.DB.Find(&book, note.BookID)
				fmt.Println(fmt.Sprintf("%s - %s ", note.Text, book.Name))
			} else {
				fmt.Println(note.Text)
			}
		},
	}
}

func (env *Env) removeDuplicatesCmd() cli.Command {
	return cli.Command{
		Name:        "deduplicate",
		Description: "Removes duplicated notes from provided book",
		ArgsUsage:   "BOOK_ID",
		Usage:       "Deduplicate notes",
		ShortName:   "d",
		Action: func(c *cli.Context) {
			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return
			}
			env.removeDuplicates(bookID)
		},
	}
}

func (env *Env) parseNotes() cli.Command {
	return cli.Command{
		Name:        "parse",
		ArgsUsage:   "FILEPATH",
		Usage:       "Parses provided file and creates notes",
		Description: "Parses file and writes books and notes",
		ShortName:   "p",
		Action: func(c *cli.Context) {
			p := c.Args().First()
			if p != "" {
				env.parseFile(p)
			}
		},
	}
}

func main() {
	// Setup SQLite database
	err := os.Mkdir(gonotesHome, os.ModePerm)
	if os.IsNotExist(err) {
		panic(err)
	}
	db, err := gorm.Open("sqlite3", fmt.Sprintf("%s/notes.db", gonotesHome))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Book{}, &Note{})
	env := &Env{DB: *db}

	// CLI app
	app := &cli.App{
		Name:                 "gonotes",
		Description:          "Simple tool to manage Kindle notes",
		EnableBashCompletion: true,
		Commands: []cli.Command{
			env.parseNotes(),
			env.listNotes(),
			env.randomNoteCmd(),
			{
				Name:      "book",
				Usage:     "Utilities to manage books",
				ShortName: "b",
				Subcommands: []cli.Command{
					env.listBooks(),
					env.deleteBook(),
					env.removeDuplicatesCmd(),
				},
			},
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
