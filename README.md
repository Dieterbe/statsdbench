Benchmark of several statsd golang clients.

note: HelperCesaro is a Grafana wrapper around Alex Cesaro's library, helper
Statsd is a Grafana wrapper around Dieterbe/statsd-go


```
$ go test -bench . -benchmem
BenchmarkAlexcesaro-8            	 5000000	       399 ns/op	       0 B/op	       0 allocs/op
BenchmarkCactus-8                	  500000	      3155 ns/op	      50 B/op	       3 allocs/op
BenchmarkCactusTimingAsDuration-8	  500000	      3283 ns/op	      82 B/op	       4 allocs/op
BenchmarkDieterbe-8              	  200000	     12004 ns/op	     352 B/op	      19 allocs/op
BenchmarkG2s-8                   	  100000	     12733 ns/op	     624 B/op	      26 allocs/op
BenchmarkHelperCesaro-8          	 3000000	       424 ns/op	       0 B/op	       0 allocs/op
BenchmarkHelperStatsd-8          	  100000	     13066 ns/op	     384 B/op	      21 allocs/op
BenchmarkQuipo-8                 	 1000000	      1851 ns/op	     400 B/op	       7 allocs/op
BenchmarkQuipoTimingAsDuration-8 	 1000000	      1447 ns/op	     192 B/op	       6 allocs/op
```
