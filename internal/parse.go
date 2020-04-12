package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"strings"
)

const line string = "=========="

// NewParseNotesCmd parses provided Kindle clippings into books and notes
func (env *Env) NewParseNotesCmd() *cli.Command {
	return &cli.Command{
		Name:        "parse",
		ArgsUsage:   "FILEPATH",
		Usage:       "Parses provided file and creates books and notes",
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

func (env *Env) parseFile(filePath string) {
	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}
	content := strings.TrimSpace(string(rawContent))
	rawNotes := strings.Split(content, line)

	for _, n := range rawNotes {
		err := env.parseAnCreateNote(n)
		if err != nil {
			log.Println(err)
		}
	}
}

func (env *Env) parseAnCreateNote(rawNote string) error {
	parts := strings.Split(strings.TrimSpace(rawNote), "\n")
	if len(parts) != 4 {
		return fmt.Errorf("failed to parse a note")
	}
	bookName := strings.TrimSpace(parts[0])
	// _ := parts[1] information about note
	text := strings.TrimSpace(parts[3])
	book := env.addBook(bookName)
	env.addNote(text, book.ID)
	return nil
}
