package main

import (
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/combo", greetall)
	http.ListenAndServe("localhost:8080", nil)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
