package handlers

import (
	"fmt"
	"net/http"
)

type Hello struct {

}

func NewHello() *Hello {
	return &Hello{}
}


func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Wold")
}
