package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "This is the beginning of a the magical Enchanted Grove adventure 🐞 ")
	})

	http.ListenAndServe(":8080", nil)

}
