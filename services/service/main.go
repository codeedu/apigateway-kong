package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("CONTENT"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
