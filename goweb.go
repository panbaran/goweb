package main

import (
	"fmt"
	"net/http"
)

//RootHandler respone for request
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0.4.0")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":3000", nil)
}
