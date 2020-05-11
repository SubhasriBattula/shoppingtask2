package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	supermarket "github.com/SubhasriBattula/shoppingtask2/supermarket"

	"github.com/gorilla/mux"
)

func getItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["item"]
	val, err := supermarket.Get(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, val)
	}
}

func postItem(w http.ResponseWriter, r *http.Request) {
	var temp map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter only item and value")
	}
	json.Unmarshal(req, &temp)
	for item, value := range temp {
		err := supermarket.Post(item, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Item entered successfully")
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	var temp map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter only item and value")
	}
	json.Unmarshal(req, &temp)
	for item, value := range temp {
		err := supermarket.Put(item, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Item are updated successfully")
		}
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["item"]
	err := supermarket.Delete(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "Items are deleted successfully")
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, supermarket.Print())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/supermarket/{item}", getItem).Methods("GET")
	router.HandleFunc("/supermarket", postItem).Methods("POST")
	router.HandleFunc("/supermarket", updateItem).Methods("PUT")
	router.HandleFunc("/supermarket/{item}", deleteItem).Methods("DELETE")
	router.HandleFunc("/supermarket", getAllItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
