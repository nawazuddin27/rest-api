package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("someone called me")
	w.WriteHeader(200)
	w.Write([]byte("hey"))
}
