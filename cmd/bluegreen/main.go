package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	bluegreen := flag.String("bluegreen", "blue", "blue or green? Default: blue")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, *bluegreen)
	})

	http.ListenAndServe(":8888", nil)
}
