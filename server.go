package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	fAddr = flag.String("addr", ":11000", "address to listen")
)

func main() {
	flag.Parse()

	words := []byte("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(words)
	})

	log.Fatal(http.ListenAndServe(*fAddr, nil))
}
