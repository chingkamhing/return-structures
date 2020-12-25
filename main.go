package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/users1", users1)
	http.HandleFunc("/users2", users2)

	http.ListenAndServe(":8888", nil)
}
