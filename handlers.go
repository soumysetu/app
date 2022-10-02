package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type customer struct {
	Name string `xml:"name" json : "name1"`
	City string `xml:"cityn" json : "city1"`
	Zip  string `xml:"zipcode" json : "zipcode1"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func greetall(w http.ResponseWriter, r *http.Request) {
	customers := []customer{
		{Name: "setu", City: "agra", Zip: "465644"},
		{Name: "soumy", City: "fdhgiri", Zip: "875644"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
