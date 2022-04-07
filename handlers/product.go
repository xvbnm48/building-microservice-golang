package handlers

import (
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

// make json encoding
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("handler GET Product")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("handler POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to marshal JSON", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

}
