package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func newQuoteFlag(dest *bool) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name: "quote", Aliases: []string{"q"}, Destination: dest, Usage: "Include book title and author",
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
			newQuoteFlag(&asQuote),
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
			env.showNotes(bookID, useIndex, asQuote)
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
			newQuoteFlag(&asQuote),
			&cli.IntFlag{Name: "length", Aliases: []string{"l"}, Destination: &lenLimit, Value: -1, Usage: "Limit note length to this value"},
		},
		Action: func(c *cli.Context) error {
			var book Book
			note := env.getRandomNote(lenLimit)
			env.DB.Find(&book, note.BookID)
			fmt.Println(getFormattedNoteText(note, book, false, asQuote))
			return nil
		},
	}
}

func (env *Env) showNotes(bookID uint, index bool, asQuote bool) {
	var notes []Note
	var book Book
	env.DB.Find(&book, bookID).Related(&notes)
	if len(notes) == 0 {
		fmt.Println("No notes :<")
		return
	}

	for _, n := range notes {
		text := getFormattedNoteText(n, book, index, asQuote)
		fmt.Printf("%s\n", text)
	}
}

func getFormattedNoteText(note Note, book Book, index bool, asQuote bool) string {
	text := note.Text
	if asQuote {
		text = fmt.Sprintf(fmt.Sprintf("\"%s\" - %s", text, book.Name))
	}
	if index {
		text = fmt.Sprintf("%d) %s", note.ID, text)
	}
	return text
}
