This little project evaluate the performance difference between a function returning an array of structure or an array of pointer of structure.

## How to run the benchmark
* invoke "make benchmark" to run the benchmark
    ```console
    go test -benchtime=10s -bench=.
    goos: linux
    goarch: amd64
    pkg: github.com/chingkamhing/return-structures
    BenchmarkArrayStruct10-4          	  396708	     29688 ns/op	    3136 B/op	      71 allocs/op
    BenchmarkArrayStruct20-4          	  201219	     59928 ns/op	    6272 B/op	     141 allocs/op
    BenchmarkArrayStruct50-4          	   81172	    148755 ns/op	   16066 B/op	     351 allocs/op
    BenchmarkArrayPointerStruct10-4   	  389179	     30171 ns/op	    3384 B/op	      76 allocs/op
    BenchmarkArrayPointerStruct20-4   	  197635	     60313 ns/op	    6776 B/op	     147 allocs/op
    BenchmarkArrayPointerStruct50-4   	   70917	    148476 ns/op	   17082 B/op	     358 allocs/op
    PASS
    ok  	github.com/chingkamhing/return-structures	75.157s
    ```

## How to run the load test
* invoke "go get -u github.com/codesenberg/bombardier" to install [bombardier](https://github.com/codesenberg/bombardier)
* invoke "./return-structures -count 20" to run the server
* invoke "make load" to perform the load test
    ```console
    bombardier -c 125 -n 10000000 -l http://localhost:8888/users-structure
    Bombarding http://localhost:8888/users-structure with 10000000 request(s) using 125 connection(s)
    10000000 / 10000000 [=====================================================================================================] 100.00% 18109/s 9m12s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec     18120.02    2893.94   27726.02
    Latency        6.90ms     3.10ms   199.13ms
    Latency Distribution
        50%     3.89ms
        75%     9.47ms
        90%    17.66ms
        95%    24.10ms
        99%    38.62ms
    HTTP codes:
        1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    30.13MB/s
    bombardier -c 125 -n 10000000 -l http://localhost:8888/users-pointer-structure
    Bombarding http://localhost:8888/users-pointer-structure with 10000000 request(s) using 125 connection(s)
    10000000 / 10000000 [======================================================================================================] 100.00% 18383/s 9m3s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec     18393.85    1952.82   26212.29
    Latency        6.80ms     2.58ms   170.52ms
    Latency Distribution
        50%     3.85ms
        75%     9.46ms
        90%    17.52ms
        95%    23.90ms
        99%    36.68ms
    HTTP codes:
        1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    30.72MB/s
    ```

## Conclusion
* benchmark suggest return array-of-structure seems tiny faster and alloc less memory
* load test confirm benchmark finding that more memory allocation affect the max and 99% latency time
* tldr: return array-of-structure has slightly better performance and less high percentile latency time
