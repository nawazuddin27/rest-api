package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	http.HandleFunc("/", greet)
	http.HandleFunc("/p", postF)

	http.HandleFunc("/p2", postF2)
	http.ListenAndServe(":8080", nil)

}

func postF(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}

func postF2(w http.ResponseWriter, r *http.Request) {
	var p person
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Name)
}

func greet(w http.ResponseWriter, r *http.Request) {

	n := r.URL.Query().Get("name")

	age := r.URL.Query().Get("age")

	a, err := strconv.Atoi(age)
	if err != nil {

		log.Fatalln(err)
	}
	w.WriteHeader(200)

	h := person{Name: n, Age: a}

	data, err := json.MarshalIndent(h, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(data)

}
