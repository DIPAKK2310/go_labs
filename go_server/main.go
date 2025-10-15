package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type Item struct {
	Id string `json:"id"`
	UserName string `json:"userName"`
}
var items = []Item{
{"1","Dipak Khare"},
{"2","Khare Dip"},
}

type Home struct {
	Id string `json:"id"`
	UserName string `json:"userName"`
}
var homes = []Home{
	{"1","Welcome to home"},
}
func getHome(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(homes)
}

func getItems(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(items)
}
func main() {
	log.Println("Server starting on 8000")

	http.HandleFunc("/items", getItems)
	http.HandleFunc("/",getHome)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal("Server failed to start ", err)
	}

}

