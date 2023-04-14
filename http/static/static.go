package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Executando")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
