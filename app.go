package main

import (
	"fmt"
	"net/http"
)

func usersStruct(w http.ResponseWriter, req *http.Request) {
	users, err := getUsersStruct(numStructs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "users 1\n")
	for i := range users {
		fmt.Fprintf(w, "%v\n", users[i])
	}
}

func usersPointerStruct(w http.ResponseWriter, req *http.Request) {
	users, err := getUsersPointerStruct(numStructs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "users 2\n")
	for i := range users {
		fmt.Fprintf(w, "%v\n", *users[i])
	}
}
