package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "HELLO, WORLD\n")
	}

	http.HandleFunc("/", helloHandler)

	log.Println("server start at port8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
