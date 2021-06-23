
#sm2-improvement

**Benchmark**
~~~~
goos: windows
goarch: amd64
pkg: sm2
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkLessThan32_P256
BenchmarkLessThan32_P256-16    	   16429	     72421 ns/op       2026 B/op	      40 allocs/op

BenchmarkLessThan32_P256SM2
BenchmarkLessThan32_P256SM2-16     13506	     88319 ns/op       2026 B/op	      40 allocs/op

BenchmarkMoreThan32_P256
BenchmarkMoreThan32_P256-16    	   15968	     74746 ns/op       2818 B/op	      46 allocs/op

BenchmarkMoreThan32_P256SM2
BenchmarkMoreThan32_P256SM2-16     13190	     90519 ns/op       2818 B/op	      46 allocs/op
>>>>>>> 同济库Benchmark Test
pkg: github.com/tjfoc/gmsm/sm2
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkSM2_LessThan32_P256SM2
BenchmarkSM2_LessThan32_P256SM2-16  777	         1525547 ns/op	   83703 B/op	    1726 allocs/op

BenchmarkSM2_MoreThan32_P256SM2
BenchmarkSM2_MoreThan32_P256SM2-16  772	         1546936 ns/op	   84076 B/op	    1725 allocs/op
