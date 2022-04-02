package main

import (
	"net/http"

	"github.com/xvbnm48/building-microservice-golang/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloWorld)

	http.HandleFunc("/goodbye", handlers.GoodBye)

	http.ListenAndServe(":8080", nil)
}
