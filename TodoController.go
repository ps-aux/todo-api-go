package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

func initTodoApi(router *mux.Router, repo *TodoRepo) {

	router.HandleFunc("/todo", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(repo.getAll())
	}).Methods("GET")

	router.HandleFunc("/todo", func(w http.ResponseWriter, req *http.Request) {
		t := Todo{}
		json.NewDecoder(req.Body).Decode(&t)
		t = repo.add(t)
		json.NewEncoder(w).Encode(t)
	}).Methods("POST")

	router.HandleFunc("/todo/{id}", func(w http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]
		repo.delete(id)
		w.WriteHeader(204)
	}).Methods("DELETE")

	router.HandleFunc("/todo/{id}", func(w http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]
		t, err := repo.get(id)

		var res interface{} = nil
		if err != nil {
			res = "Not found"
			w.WriteHeader(400)
		} else {
			res = t
		}
		json.NewEncoder(w).Encode(res)
	})

}
