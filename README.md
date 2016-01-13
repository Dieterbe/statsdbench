Benchmark of several statsd golang clients.

note: HelperCesaro is a Grafana wrapper around Alex Cesaro's library, helper
Statsd is a Grafana wrapper around Dieterbe/statsd-go


```
$ go test -bench . -benchmem -benchtime=5s
BenchmarkAlexcesaro-8            	20000000	       407 ns/op	       0 B/op	       0 allocs/op
BenchmarkCactus-8                	 2000000	      3160 ns/op	      50 B/op	       3 allocs/op
BenchmarkCactusTimingAsDuration-8	 2000000	      3307 ns/op	      82 B/op	       4 allocs/op
BenchmarkDieterbe-8              	 1000000	     12902 ns/op	     352 B/op	      19 allocs/op
BenchmarkG2s-8                   	  500000	     12937 ns/op	     624 B/op	      26 allocs/op
BenchmarkHelperCesaro-8          	20000000	       442 ns/op	       0 B/op	       0 allocs/op
BenchmarkHelperStatsd-8          	  500000	     12872 ns/op	     384 B/op	      21 allocs/op
BenchmarkQuipo-8                 	 5000000	      1849 ns/op	     400 B/op	       7 allocs/op
BenchmarkQuipoTimingAsDuration-8 	 5000000	      1437 ns/op	     192 B/op	       6 allocs/op
```
