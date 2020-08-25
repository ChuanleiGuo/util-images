package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		h, err := os.Hostname()
		if err != nil {
			io.WriteString(w, err.Error())
		}
		io.WriteString(w, h)
	}
	http.HandleFunc("/", helloHandler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
