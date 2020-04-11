package internal

import (
	"fmt"
	"github.com/urfave/cli"
)

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

// ShowNotesCmd shows notes from provided book
func (env *Env) ShowNotesCmd() cli.Command {
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

// RandomNoteCmd shows random note
func (env *Env) RandomNoteCmd() cli.Command {
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

// RemoveDuplicatesCmd performs simple deduplication of notes within single book
func (env *Env) RemoveDuplicatesCmd() cli.Command {
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

// ParseNotesCmd parses provided Kindle clippings into books and notes
func (env *Env) ParseNotesCmd() cli.Command {
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
