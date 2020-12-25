package main

import (
	"fmt"
	"net/http"
)

func users1(w http.ResponseWriter, req *http.Request) {
	users, err := getUsers1()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "users 1\n")
	for i := range users {
		fmt.Fprintf(w, "%v\n", users[i])
	}
}

func users2(w http.ResponseWriter, req *http.Request) {
	users, err := getUsers2()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "users 2\n")
	for i := range users {
		fmt.Fprintf(w, "%v\n", *users[i])
	}
}
