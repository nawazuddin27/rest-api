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

	Age int `json:"age"`
}

func main() {
	http.HandleFunc("/g", greeting)
	http.HandleFunc("/p", postf)
	http.HandleFunc("/p2", postf2)
	http.ListenAndServe(":8080", nil)
}

func postf(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func postf2(w http.ResponseWriter, r *http.Request) {
	var p person
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.Name, p.Age)
	w.Write(data)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server successfully up andrunning")
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	d, err := strconv.Atoi(age)
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(201)

	H := person{Name: name, Age: d}
	data, err := json.Marshal(H)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(data)
}
