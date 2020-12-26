.PHONY: build
build:
	go build -o return-structures *.go

.PHONY: benchmark
benchmark:
	go test -benchtime=10s -bench=.

.PHONY: profile
profile:
	go test -bench=. -cpuprofile profile_cpu.out
	go tool pprof -svg profile_cpu.out > profile_cpu.svg

.PHONY: load
load:
	bombardier -c 125 -n 10000000 -l http://localhost:8888/users-structure
	bombardier -c 125 -n 10000000 -l http://localhost:8888/users-pointer-structure

.PHONY: clean
clean:
	rm -f return-structures
