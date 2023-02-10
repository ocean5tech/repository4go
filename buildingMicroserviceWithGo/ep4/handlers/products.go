package handlers

import (
	"log"
	"net/http"

	"github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep4/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//handle update
	// goto ep4
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)   //被ToJSON替换了
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	// rw.Write(d)  //ToJSON里面包含了io.Writer，不在需要write
}
