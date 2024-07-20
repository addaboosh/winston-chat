package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (rw http.ResponseWriter, r *http.Request){
		fmt.Println("Hello World")
	})
	http.ListenAndServe(":8123", nil)


}
