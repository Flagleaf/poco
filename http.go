package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
	})
	http.ListenAndServe(":9090", nil)
}
