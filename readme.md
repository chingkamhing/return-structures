This little project evaluate the performance difference between a function returning an array of structure or an array of pointer of structure.

## How to run the benchmark
* invoke "make benchmark" to run the benchmark
    ```console
    go test -bench=.
    goos: linux
    goarch: amd64
    pkg: github.com/chingkamhing/test/return-structures
    BenchmarkArrayStruct-4                     35564             31852 ns/op            3984 B/op         35 allocs/op
    BenchmarkArrayPointerStruct-4              35866             32691 ns/op            4488 B/op         41 allocs/op
    PASS
    ok      github.com/chingkamhing/test/return-structures  2.999s
    ```

## How to run the load test
* invoke "./return-structures" to run the server
* invoke "make load" to perform the load test
    ```console
    kamching@kamching-Aspire-E5-571G:~$ bombardier -c 125 -n 10000000 -l http://localhost:8888/users1
    Bombarding http://localhost:8888/users1 with 10000000 request(s) using 125 connection(s)
    10000000 / 10000000 [======================================================================================================] 100.00% 27365/s 6m5s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec     27371.57    1907.01   41466.18
    Latency        4.56ms     1.35ms    75.61ms
    Latency Distribution
        50%     2.91ms
        75%     6.38ms
        90%    11.28ms
        95%    14.91ms
        99%    23.26ms
    HTTP codes:
        1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    37.04MB/s
    kamching@kamching-Aspire-E5-571G:~$ bombardier -c 125 -n 10000000 -l http://localhost:8888/users2
    Bombarding http://localhost:8888/users2 with 10000000 request(s) using 125 connection(s)
    10000000 / 10000000 [=====================================================================================================] 100.00% 25997/s 6m24s
    Done!
    Statistics        Avg      Stdev        Max
    Reqs/sec     26006.63    3839.56   37901.93
    Latency        4.80ms     1.86ms   153.43ms
    Latency Distribution
        50%     2.93ms
        75%     6.59ms
        90%    11.98ms
        95%    16.01ms
        99%    25.86ms
    HTTP codes:
        1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
        others - 0
    Throughput:    35.19MB/s
    ```

## Conclusion
* benchmark suggest return array-of-structure seems tiny faster and alloc less memory
* load test confirm benchmark finding that more memory allocation affect the max and 99% latency time
* tldr: return array-of-structure has slightly better performance and less high percentile latency time
