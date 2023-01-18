package main

import (
	"fmt"
	"net/http"
)

//RootHandler respone for request
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Your Go application has been successfully deployed on Kubernetes.")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":3000", nil)
}
