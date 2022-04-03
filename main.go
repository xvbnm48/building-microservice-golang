package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xvbnm48/building-microservice-golang/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api : ", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)

	http.ListenAndServe(":8080", sm)
}
