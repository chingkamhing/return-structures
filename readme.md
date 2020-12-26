This little project evaluate the performance difference between a function returning an array of structure or an array of pointer of structure.

## How to run the benchmark
* invoke "make benchmark" to run the benchmark
    ```console
    go test -benchtime=10s -bench=.
    goos: linux
    goarch: amd64
    pkg: github.com/chingkamhing/return-structures
    BenchmarkArrayStruct1-4           	 3890893	      3107 ns/op	     320 B/op	       8 allocs/op
    BenchmarkArrayStruct10-4          	  400992	     29599 ns/op	    3136 B/op	      71 allocs/op
    BenchmarkArrayStruct20-4          	  201361	     60596 ns/op	    6272 B/op	     141 allocs/op
    BenchmarkArrayStruct50-4          	   80115	    156941 ns/op	   16066 B/op	     351 allocs/op
    BenchmarkArrayPointerStruct1-4    	 3784086	      3180 ns/op	     328 B/op	       9 allocs/op
    BenchmarkArrayPointerStruct10-4   	  395434	     30279 ns/op	    3384 B/op	      76 allocs/op
    BenchmarkArrayPointerStruct20-4   	  198122	     61174 ns/op	    6776 B/op	     147 allocs/op
    BenchmarkArrayPointerStruct50-4   	   79572	    149561 ns/op	   17082 B/op	     358 allocs/op
    PASS
    ok  	github.com/chingkamhing/return-structures	107.967s
    ```

## How to run the load test
* invoke "go get -u github.com/codesenberg/bombardier" to install [bombardier](https://github.com/codesenberg/bombardier)
* invoke "./return-structures -count 20" to run the server
* invoke "make load" to perform the load test
    ```console
    bombardier -c 125 -n 10000000 -l http://localhost:8888/users-structure
    Bombarding http://localhost:8888/users-structure with 10000000 request(s) using 125 connection(s)
    10000000 / 10000000 [=====================================================================================================] 100.00% 18541/s 8m59s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec     18553.72    2070.39   26958.62
    Latency        6.74ms     2.59ms   178.18ms
    Latency Distribution
        50%     3.79ms
        75%     9.30ms
        90%    17.42ms
        95%    23.72ms
        99%    37.07ms
    HTTP codes:
        1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    30.85MB/s
    bombardier -c 125 -n 10000000 -l http://localhost:8888/users-pointer-structure
    Bombarding http://localhost:8888/users-pointer-structure with 10000000 request(s) using 125 connection(s)
    10000000 / 10000000 [=====================================================================================================] 100.00% 18098/s 9m12s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec     18109.22    2272.68   26363.56
    Latency        6.90ms     2.83ms   212.26ms
    Latency Distribution
        50%     3.80ms
        75%     9.52ms
        90%    17.94ms
        95%    24.66ms
        99%    38.06ms
    HTTP codes:
        1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    30.25MB/s
    ```

## Conclusion
* benchmark suggest return array-of-structure seems tiny faster and alloc less memory
* load test confirm benchmark finding that more memory allocation affect the max and 99% latency time
* tldr: return array-of-structure has slightly better performance and less high percentile latency time
