package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	firstName string
	lastName  string
	age       int
	company   []string
}

func main() {
	fmt.Println("My Website...")
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/detail", DetailPage)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintln(w, "Hello http page")
}

func DetailPage(w http.ResponseWriter, r *http.Request) {
	var detailInfo = Info{
		firstName: "Le Anh",
		lastName:  "Quang",
		age:       23,
		company:   []string{"MCG Corp", "CMC Global"},
	}
	json.NewEncoder(w).Encode(detailInfo)
}
