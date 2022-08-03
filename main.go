package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("hogehoge")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your url path is %s", r.URL.Path[1:])
}
