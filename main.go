package main

import (
	"flag"
	"net/http"
)

// number of structures to response
var numStructs int

func init() {
	flag.IntVar(&numStructs, "count", 10, "number of structures to response")
}

func main() {
	// parse flags
	flag.Parse()
	// prepare endpoints
	http.HandleFunc("/users-structure", usersStruct)
	http.HandleFunc("/users-pointer-structure", usersPointerStruct)
	// start http server
	http.ListenAndServe(":8888", nil)
}
