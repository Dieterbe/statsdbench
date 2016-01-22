Benchmark of several statsd golang clients.

note: HelperCesaro is a raintank wrapper around Alex Cesaro's library
DieterbeRaw has the client write byte arrays which may not always be feasible.

```
$ go test -bench . -benchmem
BenchmarkAlexcesaro-8            	 3000000	       408 ns/op	       0 B/op	       0 allocs/op
BenchmarkCactus-8                	  500000	      3143 ns/op	      50 B/op	       3 allocs/op
BenchmarkCactusTimingAsDuration-8	  500000	      3323 ns/op	      82 B/op	       4 allocs/op
BenchmarkDieterbe-8              	  200000	     11672 ns/op	     352 B/op	      19 allocs/op
BenchmarkDieterbeRaw-8           	  200000	      7473 ns/op	       0 B/op	       0 allocs/op
BenchmarkG2s-8                   	  200000	     11704 ns/op	     624 B/op	      26 allocs/op
BenchmarkHelperCesaro-8          	 3000000	       434 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuipo-8                 	 1000000	      1841 ns/op	     400 B/op	       7 allocs/op
BenchmarkQuipoTimingAsDuration-8 	 1000000	      1427 ns/op	     192 B/op	       6 allocs/op
```
