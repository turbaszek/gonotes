package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Books endpoint handler
func (env *Env) Books(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var book Book
	var books []Book

	switch r.Method {
	case "GET":
		// Get single book
		if id := vars["id"]; id != "" {
			env.Db.Find(&book, id)
			if book.ID != 0 {
				json.NewEncoder(w).Encode(book)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
			// Get all books
		} else {
			env.Db.Find(&books)
			json.NewEncoder(w).Encode(books)
			w.WriteHeader(http.StatusOK)
		}
	case "POST":
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()
		err := d.Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if env.Db.NewRecord(&book) {
			env.Db.Create(&book)
			w.WriteHeader(http.StatusCreated)
		}
	}
}

// Notes endpoint handler
func (env *Env) Notes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		var notes []Note
		bookID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "No book id", http.StatusBadRequest)
			return
		}
		env.Db.Find(&notes, Note{BookID: uint(bookID)})
		json.NewEncoder(w).Encode(notes)
		w.WriteHeader(http.StatusOK)
	case "POST":
		var note Note
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()
		err := d.Decode(&note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if env.Db.NewRecord(&note) {
			env.Db.Create(&note)
			w.WriteHeader(http.StatusCreated)
		}
	case "DELETE":
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var note Note
		env.Db.Delete(&note, id)
		w.WriteHeader(http.StatusOK)
	}

}
