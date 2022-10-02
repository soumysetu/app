package main

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/combo", greetall)
	
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
