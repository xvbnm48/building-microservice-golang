package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xvbnm48/building-microservice-golang/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}

	w.Write(d)
}
