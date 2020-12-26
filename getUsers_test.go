package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkArrayStruct1(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersStruct(1)
		for i := range users {
			fmt.Fprintf(file, "%v\n", users[i])
		}
	}
}

func BenchmarkArrayStruct10(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersStruct(10)
		for i := range users {
			fmt.Fprintf(file, "%v\n", users[i])
		}
	}
}

func BenchmarkArrayStruct20(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersStruct(20)
		for i := range users {
			fmt.Fprintf(file, "%v\n", users[i])
		}
	}
}

func BenchmarkArrayStruct50(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersStruct(50)
		for i := range users {
			fmt.Fprintf(file, "%v\n", users[i])
		}
	}
}

func BenchmarkArrayPointerStruct1(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersPointerStruct(1)
		for i := range users {
			fmt.Fprintf(file, "%v\n", *users[i])
		}
	}
}

func BenchmarkArrayPointerStruct10(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersPointerStruct(10)
		for i := range users {
			fmt.Fprintf(file, "%v\n", *users[i])
		}
	}
}

func BenchmarkArrayPointerStruct20(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersPointerStruct(20)
		for i := range users {
			fmt.Fprintf(file, "%v\n", *users[i])
		}
	}
}

func BenchmarkArrayPointerStruct50(b *testing.B) {
	file, err := os.Open("/dev/null")
	if err != nil {
		return
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		users, _ := getUsersPointerStruct(50)
		for i := range users {
			fmt.Fprintf(file, "%v\n", *users[i])
		}
	}
}
