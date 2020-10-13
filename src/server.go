package main

import (
	"fast-topic-matching/srcl/matching"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var matcher = NewCSTrieMatcher()

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/subscribe", subscribe)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fascriber := Fascriber{WebSocket, "fas123"}
	_, err := matcher.Subscribe("vehicle/motional0001/state", fascriber)
	if err == nil {
		fascriber.ConfirmSubscription()
	}
}
