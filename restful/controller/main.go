package main

import (
	"github.com/gorilla/mux"
	"learn-go/restful/service"
	"log"
	"net/http"
)


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/book/{id}",service.GetBookById).Methods("GET")



	log.Fatal(http.ListenAndServe(":8888",router))
}
