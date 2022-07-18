package main

import (
	"encoding/json"
	"log"
	"net/http"

	db "assistanceTool.com/panelbackend/db"
	dl "assistanceTool.com/panelbackend/service"
	"github.com/gorilla/mux"
)

func main() {

	log.Printf("Starting backend...")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode("Ok")
		db.ErrorCheck(err)
	})

	myRouter.HandleFunc("/register",dl.RegisterHandler).Methods("POST","OPTIONS")
	myRouter.HandleFunc("/login",dl.LoginHandler).Methods("POST","OPTIONS")
	myRouter.HandleFunc("/contact",dl.ContactHandler).Methods("GET","OPTIONS")

    log.Fatalln(http.ListenAndServe(":5010", myRouter))
}

