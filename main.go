package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/george4joseph/go_library_blockchain/controllers"
	"github.com/george4joseph/go_library_blockchain/models"
)



var Blockch *models.BlockChain




func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book", controllers.newBook).Methods("POST")
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/issue", writeBlock).Methods("POST")

	log.Println("Listening to requests... port is 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
