package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func shortName(text string, length int) string {
	// Changing : to - to because part after : is treated as description in cli etc
	text = strings.ReplaceAll(text, ":", "-")
	if len(text) > length {
		return fmt.Sprintf("%s...", text[:length-3])
	}
	return text
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

// ShowNotesCmd shows notes from provided book
func (env *Env) ShowNotesCmd() *cli.Command {
	var useIndex bool
	var books []Book
	env.DB.Order("id desc").Find(&books)

	return &cli.Command{
		Name:        "notes",
		Description: "Lists all notes from provided book",
		ArgsUsage:   "BOOK_ID",
		Usage:       "List notes",
		Aliases:     []string{"n"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "index, id", Destination: &useIndex, Usage: "If set notes id will be included"},
		},
		Action: func(c *cli.Context) error {
			// Autocomplete provides {{Book.ID}} - {{Book.Name}}
			args := strings.Split(c.Args().First(), "-")
			bookIDString := args[len(args)-1]
			bookID, err := strToUint(bookIDString)
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return err
			}
			env.showNotes(bookID, useIndex)
			return nil
		},
		BashComplete: func(c *cli.Context) {
			// This will complete if no args are passed
			if c.NArg() > 0 {
				return
			}
			for _, b := range books {
				fmt.Println(fmt.Sprintf("%s-%d", shortName(b.Name, 70), b.ID))
			}
		},
	}
}

// RandomNoteCmd shows random note
func (env *Env) RandomNoteCmd() *cli.Command {
	var asQuote bool
	var lenLimit int
	return &cli.Command{
		Name:        "random",
		Description: "Shows a random note",
		Usage:       "Shows a random note",
		Aliases:     []string{"r"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "quote, q", Destination: &asQuote, Usage: "Include book title and author"},
			&cli.IntFlag{Name: "length, l", Destination: &lenLimit, Value: -1, Usage: "Limit note length to this value"},
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

// RemoveDuplicatesCmd performs simple deduplication of notes within single book
func (env *Env) RemoveDuplicatesCmd() *cli.Command {
	return &cli.Command{
		Name:        "deduplicate",
		Description: "Removes duplicated notes from provided book",
		ArgsUsage:   "BOOK_ID",
		Usage:       "Deduplicate notes",
		Aliases:     []string{"d"},
		Action: func(c *cli.Context) error {
			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return err
			}
			env.removeDuplicates(bookID)
			return nil
		},
	}
}

// ParseNotesCmd parses provided Kindle clippings into books and notes
func (env *Env) ParseNotesCmd() *cli.Command {
	return &cli.Command{
		Name:        "parse",
		ArgsUsage:   "FILEPATH",
		Usage:       "Parses provided file and creates notes",
		Description: "Parses file and writes books and notes",
		Aliases:     []string{"p"},
		Action: func(c *cli.Context) error {
			p := c.Args().First()
			if p != "" {
				env.parseFile(p)
			}
			return nil
		},
	}
}
