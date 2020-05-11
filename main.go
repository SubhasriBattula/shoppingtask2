package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	super "github.com/SubhasriBattula/shoppingtask2/super"

	"github.com/gorilla/mux"
)

func itemGet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["item"]
	val, err := super.Get(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, val)
	}
}

func itemPost(w http.ResponseWriter, r *http.Request) {
	var temp map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data in correct format")
	}
	json.Unmarshal(req, &temp)
	for item, value := range temp {
		err := super.Post(item, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Item entered successfully")
		}
	}
}
func itemUpdate(w http.ResponseWriter, r *http.Request) {
	var temp map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data in correct format")
	}
	json.Unmarshal(req, &temp)
	for item, value := range temp {
		err := super.Put(item, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Item updated successfully")
		}
	}
}

func itemDelete(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["item"]
	err := super.Delete(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "Items deleted successfully")
	}
}

func itemGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, super.Print())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/supermarket/{item}", itemGet).Methods("GET")
	router.HandleFunc("/supermarket", itemPost).Methods("POST")
	router.HandleFunc("/supermarket", itemUpdate).Methods("PUT")
	router.HandleFunc("/supermarket/{item}", itemDelete).Methods("DELETE")
	router.HandleFunc("/supermarket", itemGetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
