package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func newQuoteFlag(dest bool) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name: "quote", Aliases: []string{"q"}, Destination: &dest, Usage: "Include book title and author",
	}
}

// NewShowNotesCmd shows notes from provided book
func (env *Env) NewShowNotesCmd() *cli.Command {
	var useIndex bool
	var asQuote bool
	return &cli.Command{
		Name:        "notes",
		Description: "Lists all notes from provided book. Supports autocomplete!",
		ArgsUsage:   "BOOK_ID",
		Usage:       "List notes",
		Aliases:     []string{"n"},
		Flags: []cli.Flag{
			newQuoteFlag(asQuote),
			&cli.BoolFlag{Name: "index, id", Destination: &useIndex, Usage: "If set notes id will be included"},
		},
		Action: func(c *cli.Context) error {
			// Autocomplete provides {{Book.ID}}:{{Book.Name}} in case of bash
			args := strings.Split(c.Args().First(), ":")
			bookIDString := args[0]
			bookID, err := strToUint(bookIDString)
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return err
			}
			env.showNotes(bookID, useIndex)
			return nil
		},
		BashComplete: env.booksComplete,
	}
}

// NewRandomNoteCmd shows random note
func (env *Env) NewRandomNoteCmd() *cli.Command {
	var asQuote bool
	var lenLimit int
	return &cli.Command{
		Name:        "random",
		Description: "Shows a random note",
		Usage:       "Shows a random note",
		Aliases:     []string{"r"},
		Flags: []cli.Flag{
			newQuoteFlag(asQuote),
			&cli.IntFlag{Name: "length", Aliases: []string{"l"}, Destination: &lenLimit, Value: -1, Usage: "Limit note length to this value"},
		},
		Action: func(c *cli.Context) error {
			note := env.getRandomNote(lenLimit)
			if asQuote {
				book := Book{}
				env.DB.Find(&book, note.BookID)
				fmt.Println(fmt.Sprintf("%s - %s ", note.Text, book.Name))
			} else {
				fmt.Println(note.Text)
			}
			return nil
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
			fmt.Printf("%d) %s\n", n.ID, n.Text)
		} else {
			fmt.Printf("%s\n", n.Text)
		}
	}
}
