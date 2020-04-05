package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

var home, _ = os.UserHomeDir()
var gonotesHome = fmt.Sprintf("%s/gonotes", home)

const ls = "ls"
const rm = "rm"
const cat = "cat"

func strToUint(n string) (uint, error) {
	v, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func (env *Env) listBooks() cli.Command {
	return cli.Command{
		Name:  ls,
		Usage: "List books",
		Action: func(c *cli.Context) {
			var books []Book
			env.DB.Find(&books)
			if len(books) == 0 {
				fmt.Println("It seems you have no books yet. Try to parse clippings from your Kindle.")
			}
			for i := 0; i < len(books); i++ {
				b := books[i]
				fmt.Println(fmt.Sprintf("%d | %s", b.ID, b.Name))
			}
		},
	}
}

func (env *Env) listNotes() cli.Command {
	return cli.Command{
		Name:  cat,
		Usage: "List notes",
		Action: func(c *cli.Context) {
			var notes []Note

			bookID, err := strToUint(c.Args().First())
			if err != nil {
				fmt.Println("ID of a book have to be a integer")
				return
			}
			env.DB.Where(Note{bookID: bookID}).Find(&notes)
			if len(notes) == 0 {
				fmt.Println("No notes :<")
				return
			}
			for i := 0; i < len(notes); i++ {
				n := notes[i]
				fmt.Println(fmt.Sprintf("%d | %s\n", n.ID, n.Text))
			}
		},
	}
}

func (env *Env) deleteBook() cli.Command {
	return cli.Command{
		Name:  rm,
		Usage: "Delete book id",
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

func (env *Env) readNotes() cli.Command {
	return cli.Command{
		Name:  "parse",
		Usage: "Parses provided file",
		Action: func(c *cli.Context) {
			env.parseFile(c.Args().First())
		},
	}
}

func main() {
	// Connect to DB
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

	app := &cli.App{
		Name:  "GoNotes",
		Usage: "Simple cli tool to manage Kindle notes",
		Commands: []cli.Command{
			env.readNotes(),
			{
				Name:  "note",
				Usage: "Notes related operations",
				Subcommands: []cli.Command{
					env.listNotes(),
				},
			},
			{
				Name:  "book",
				Usage: "Books related operations",
				Subcommands: []cli.Command{
					env.listBooks(),
					env.deleteBook(),
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
