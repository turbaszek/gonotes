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
	content := strings.TrimSpace(string(rawContent[:]))
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
	l := len(parts)
	if l < 4 {
		return nil // skip when note's text is empty
	}
	if l > 4 {
		return fmt.Errorf("failed to parse a note. Had %d lines, expected 4", l)
	}
	bookName := trimText(parts[0])
	// _ := parts[1] information about note
	text := trimText(parts[3])
	book := env.addBook(bookName)
	env.addNote(text, book.ID)
	return nil
}

// It seems that sometimes book titles starts with strange symbols that
// may result in duplication of a book. It's hard to spot because reading
// it from database shows same string.
func isStrangeRune(r rune) bool {
	excluded := []rune{
		65279, // U+00EF Latin small letter i, diaeresis
	}
	for _, x := range excluded {
		if r == x {
			return true
		}
	}
	return false
}
func trimText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.TrimLeftFunc(text, isStrangeRune)
	return text
}
