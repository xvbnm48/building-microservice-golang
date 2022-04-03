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
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":8080", sm)
}
