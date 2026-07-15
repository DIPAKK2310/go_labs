package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}
type Result struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

var items = []Item{
	{"1", "Dipak Khare"},
	{"2", "Khare Dipak"},
}
var result = []Result{
	{"1","Pass"},
	{"2","fail"},
}

type Home struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

var homes = []Home{
	{"1", "Welcome to home"},
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(homes)
}
func getResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(homes)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func main() {
	log.Println("Server starting on https://www.localhost:8000")

	http.HandleFunc("/items", getItems)
	http.HandleFunc("/", getHome)
	http.HandleFunc("/result", getResult)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Server failed to start ", err)
	}

}
