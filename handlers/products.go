package handlers

import (
	"log"
	"net/http"

	"github.com/yudafatah/product-api-learngolang/data"
)

//Products structure
type Products struct {
	l *log.Logger
}

//NewProducts return product list
func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		p.getProducts(w,r)
		return
	}

	// catch all HTTP
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request){
	lp := data.GetProduct()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
}