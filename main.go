package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ping pong")
	})

	http.ListenAndServe(":8080", nil)
}