package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "oopss", http.StatusBadRequest)

			return
		}
		// log.Printf("Data %s\n", d)
		fmt.Fprintf(w, "Hello World %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":8080", nil)
}
