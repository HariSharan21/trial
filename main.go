package main

import (
	"fmt"
	"net/http"
	"trial/api"
)

func main() {
	fmt.Println("api started")
	http.HandleFunc("/sum", api.SumHandler)
	http.ListenAndServe(":8080", nil)
}
