package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/george4joseph/go_library_blockchain/models"
)

func GetBlockChain(blockc *models.BlockChain) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {

		jbytes, err := json.MarshalIndent(blockc.Blocks_, "", " ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		io.WriteString(w, string(jbytes))
	}

}
