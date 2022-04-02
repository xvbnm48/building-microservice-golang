package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "oopss", http.StatusBadRequest)

		return
	}
	// log.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Hello World %s", d)
}

func GoodBye(w http.ResponseWriter, r *http.Request) {
	log.Println("good bye")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "oopps", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Good Bye %s", d)
}
