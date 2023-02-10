package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l} // 一定要写l，否则无法正确注入
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// ServeHTTP的大小写不能错，是interface名

	/* 	type Handler interface {
		   ServeHTTP(ResponseWriter, *Request)
	} */

	h.l.Println("Hello World.")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s\n", d)

	fmt.Fprintf(rw, "Hello %s\n", d)
}
