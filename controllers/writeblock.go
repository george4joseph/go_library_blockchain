package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/george4joseph/go_library_blockchain/models"
)


func WriteBlock(bc *models.BlockChain) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){

		var bookissue models.BookIssue

		if err := json.NewDecoder(r.Body).Decode(&bookissue); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Could not write book issue: %v", err)
			w.Write([]byte("Could not write book issue"))

		}

	bc.AddBlock(bookissue)

	resp, err := json.MarshalIndent(bookissue, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	}
}
