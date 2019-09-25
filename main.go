package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-template-api/controllers"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/examples", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/examples", controllers.GetAllExamples).Methods("GET")
	router.HandleFunc("/examples/{id}", controllers.GetOneEvent).Methods("GET")
	router.HandleFunc("/examples/{id}", controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/examples/{id}", controllers.DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}