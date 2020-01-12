package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	portStr := os.Getenv("SERVER_PORT")

	if len(portStr) == 0 {
		portStr = "8080"
	}

	port, err := strconv.ParseInt(portStr, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting at port %d \n", port)
	//Write(Todo{
	//	Id:   "123",
	//	Name: "blabla",
	//	Due:  time.Now()})

	initTodoApi(router, &TodoRepo{})

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello from Go TODO api")
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))

}
