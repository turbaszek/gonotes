package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

// NewListBooksCmd list all available books
func (env *Env) NewListBooksCmd() *cli.Command {
	return &cli.Command{
		Name:        "ls",
		ArgsUsage:   "",
		Description: "Lists all books",
		Usage:       "Lists books",
		Action: func(c *cli.Context) error {
			var books []Book
			env.DB.Find(&books)
			if len(books) == 0 {
				fmt.Println("It seems you have no books yet. Try to parse clippings from your Kindle.")
				return nil
			}
			for _, b := range books {
				fmt.Printf("%d - %s\n", b.ID, b.Name)
			}
			return nil
		},
	}
}

// NewDeleteBookCmd deletes selected book by ID
func (env *Env) NewDeleteBookCmd() *cli.Command {
	return &cli.Command{
		Name:        "rm",
		Description: "Deletes books by ID. You can provide multiple IDs. Supports autocomplete!",
		ArgsUsage:   "BOOK_ID_1 BOOK_ID_2 ...",
		Usage:       "Deletes books",
		Action: func(c *cli.Context) error {
			for _, arg := range c.Args().Slice() {
				bookID, err := strToUint(arg)
				if err != nil {
					fmt.Println("ID of a book have to be a integer")
					return err
				}
				env.removeBook(bookID)
			}
			return nil
		},
		BashComplete: func(context *cli.Context) {
			// This complete works with multiple arguments
			var books []Book
			env.DB.Order("id desc").Find(&books)
			for _, b := range books {
				fmt.Printf("%d:%s\n", b.ID, shortName(b.Name, 70))
			}
		},
	}
}

// NewRemoveDuplicatesCmd performs simple deduplication of notes within single book
func (env *Env) NewRemoveDuplicatesCmd() *cli.Command {
	var all bool
	return &cli.Command{
		Name:        "deduplicate",
		Description: "Removes duplicated notes from provided book. Supports autocomplete!",
		ArgsUsage:   "BOOK_ID",
		Usage:       "Deduplicate notes",
		Aliases:     []string{"d"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Destination: &all, Usage: "Deduplicate all books"},
		},
		Action: func(c *cli.Context) error {
			if all {
				var books []Book
				env.DB.Find(&books)
				for _, b := range books {
					env.removeDuplicates(b.ID)
				}
				return nil
			}
			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return err
			}
			env.removeDuplicates(bookID)
			return nil
		},
		BashComplete: env.booksComplete,
	}
}

func shortName(text string, length int) string {
	if len(text) > length {
		return fmt.Sprintf("%s...", text[:length-3])
	}
	return text
}

func (env *Env) booksComplete(c *cli.Context) {
	var books []Book
	env.DB.Order("id desc").Find(&books)
	// This will complete if no args are passed
	if c.NArg() > 0 {
		return
	}
	for _, b := range books {
		fmt.Printf("%d:%s\n", b.ID, shortName(b.Name, 70))
	}
}

func (env *Env) removeDuplicates(bookID uint) {
	var notes []Note
	env.DB.Where("book_id == ?", bookID).Find(&notes)

	i := 0
	for i < len(notes) {
		j := i + 1
		n := notes[i]
		for j < len(notes) {
			if strings.Contains(n.Text, notes[j].Text) {
				env.removeNote(notes[j].ID)
				notes[j] = notes[len(notes)-1]
				notes = notes[:len(notes)-1]
			} else {
				if strings.Contains(notes[j].Text, n.Text) {
					env.removeNote(notes[i].ID)
					notes[i] = notes[j]
					notes[j] = notes[len(notes)-1]
					notes = notes[:len(notes)-1]
				}
			}
			j++
		}
		i++
	}
}
