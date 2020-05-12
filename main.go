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
	productname := mux.Vars(r)["listitems"]
	val, err := supermarket.GetProduct(productname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintln(w, val)
		w.WriteHeader(http.StatusOK)
	}
}

func postItem(w http.ResponseWriter, r *http.Request) {
	var list map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter only item and value")
	}
	json.Unmarshal(req, &list)
	for listitems, value := range list {
		err := supermarket.PostProduct(listitems, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Items entered successfully")
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	var list map[string]interface{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter only item and value")
	}
	json.Unmarshal(req, &list)
	for listitems, value := range list {
		err := supermarket.PutProduct(listitems, value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Items are updated successfully")
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	productname := mux.Vars(r)["listitems"]
	err := supermarket.DeleteProduct(productname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "Items are deleted successfully")
		w.WriteHeader(http.StatusAccepted)
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, supermarket.PrintProduct())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/supermarket/{listitems}", getItem).Methods("GET")
	router.HandleFunc("/supermarket", postItem).Methods("POST")
	router.HandleFunc("/supermarket", updateItem).Methods("PUT")
	router.HandleFunc("/supermarket/{listitems}", deleteItem).Methods("DELETE")
	router.HandleFunc("/supermarket", getAllItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
