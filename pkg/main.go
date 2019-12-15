package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nuclearpinguin/gonotes/pkg/app"
	"net/http"
)

func main() {
	// Connect to DB
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&app.Book{}, &app.Note{})
	env := &app.Env{Db: db}

	// API endpoints
	r := mux.NewRouter()
	r.HandleFunc("/books/", env.Books).Methods("POST", "GET")
	r.HandleFunc("/books/{id:[0-9]*}", env.Books).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}/notes", env.Notes).Methods("GET", "POST")
	r.HandleFunc("/notes/{id:[0-9]*}", env.Notes).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
