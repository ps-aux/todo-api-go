package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	log.Println("starting")

	//Write(Todo{
	//	Id:   "123",
	//	Name: "blabla",
	//	Due:  time.Now()})

	initTodoApi(router, &TodoRepo{})

	log.Fatal(http.ListenAndServe(":8080", router))

}
