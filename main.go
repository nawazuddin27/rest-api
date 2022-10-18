package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type person struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/s", mygreetings)
	http.ListenAndServe(":8080", nil)
}

func mygreetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("succesfully run the server")
	n := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	d, err := strconv.Atoi(age)
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(201)

	H := person{Name: n, Age: d}
	data, err := json.Marshal(H)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(data)

}
