package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/george4joseph/go_library_blockchain/models"
)

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var bookissue models.BookIssue

	if err := json.NewDecoder(r.Body).Decode(&bookissue); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not write book issue: %v", err)
		w.Write([]byte("Could not write book issue"))

	}

	models.BlockChain.AddBlock(data bookissue)

}




