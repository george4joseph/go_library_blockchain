package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/george4joseph/go_library_blockchain/models"
)

func newBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create book: %v", err)
		w.Write([]byte("Could not create book:"))
		return

	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(book,""," ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal payload: %v", err)
		w.Write([]byte("Could not save book: "))

	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	

}
