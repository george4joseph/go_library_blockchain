package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	"github.com/george4joseph/go_library_blockchain/controllers"
	"github.com/george4joseph/go_library_blockchain/models"
)

var blockch *models.BlockChain

func main() {

	blockch = controllers.NewBlockChain()
	r := mux.NewRouter()
	r.HandleFunc("/book", controllers.NewBook).Methods("POST")
	r.HandleFunc("/", controllers.GetBlockChain(blockch)).Methods("GET")
	r.HandleFunc("/issue", controllers.WriteBlock(blockch)).Methods("POST")

	go func() {
		for _, block := range blockch.Blocks_ {
			fmt.Printf("Prev hash: %x\n", block.PreHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data:  %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Listening to requests... port is 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
