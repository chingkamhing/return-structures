package main

import (
	"fmt"
)

// User define the user fields
type User struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	Birthday  int64
}

// getUsersStruct simulate to get a list of users information
func getUsersStruct(count int) ([]User, error) {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			FirstName: fmt.Sprintf("FirstName 1%03d", i),
			LastName:  fmt.Sprintf("LastName 1%03d", i),
			Email:     fmt.Sprintf("Email 1%03d", i),
			Phone:     fmt.Sprintf("Phone 1%03d", i),
			Address:   fmt.Sprintf("Address 1%03d", i),
			Birthday:  int64(1609459100 + i),
		}
	}
	return users, nil
}

// getUsersPointerStruct simulate to get a list of users information
func getUsersPointerStruct(count int) ([]*User, error) {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			FirstName: fmt.Sprintf("FirstName 2%03d", i),
			LastName:  fmt.Sprintf("LastName 2%03d", i),
			Email:     fmt.Sprintf("Email 2%03d", i),
			Phone:     fmt.Sprintf("Phone 2%03d", i),
			Address:   fmt.Sprintf("Address 2%03d", i),
			Birthday:  int64(1609459200 + i),
		}
	}
	var usersPtr []*User
	for i := range users {
		usersPtr = append(usersPtr, &users[i])
	}
	return usersPtr, nil
}
