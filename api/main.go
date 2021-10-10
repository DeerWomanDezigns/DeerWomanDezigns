package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func orderPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "orderPage!")
	fmt.Println("Endpoint Hit: orderPage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/order", orderPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
