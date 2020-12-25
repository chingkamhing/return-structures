package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkArrayStruct(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsers1()
		for i := range users {
			fmt.Fprintf(file, "%v\n", users[i])
		}
	}
}

func BenchmarkArrayPointerStruct(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsers2()
		for i := range users {
			fmt.Fprintf(file, "%v\n", *users[i])
		}
	}
}
