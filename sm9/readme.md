## sm9-improvement

根据论文《一种基于FPGA的SM9快速实现方法》、《一种SM9数字签名及验证算法的快速实现方法》进行部分优化，同时利用https://github.com/xlcetc/cryptogm提供的sm9曲线包提高性能。原曲线包为https://github.com/cloudflare/bn256

后期若将该项目库中的sm3一同使用或许效果更佳

Benchmark

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkSign-16             	     643	   1834228 ns/op	   49728 B/op	     444 allocs/op
BenchmarkNewSign-16          	    1284	    920666 ns/op	    2495 B/op	      23 allocs/op
BenchmarkNewVerify-16        	     574	   2093523 ns/op	   50040 B/op	     443 allocs/op
BenchmarkNewSignVerify-16    	     393	   3017218 ns/op	   52656 B/op	     467 allocs/op
BenchmarkNewSignLen/BenchmarkNewSignLen20-16         	    1294	    917279 ns/op	    2460 B/op	      23 allocs/op
BenchmarkNewSignLen/BenchmarkNewSignLen40-16         	    1262	    921380 ns/op	    2524 B/op	      23 allocs/op
BenchmarkNewSignLen/BenchmarkNewSignLen80-16         	    1299	    916463 ns/op	    2588 B/op	      23 allocs/op
BenchmarkNewSignLen/BenchmarkNewSignLen160-16        	    1297	    921074 ns/op	    2780 B/op	      23 allocs/op
BenchmarkNewSignLen/BenchmarkNewSignLen320-16        	    1293	    919335 ns/op	    3166 B/op	      23 allocs/op
BenchmarkNewVerifyLen/BenchmarkNewVerifyLen20-16     	     580	   2060416 ns/op	   50040 B/op	     443 allocs/op
BenchmarkNewVerifyLen/BenchmarkNewVerifyLen40-16     	     583	   2057580 ns/op	   50104 B/op	     443 allocs/op
BenchmarkNewVerifyLen/BenchmarkNewVerifyLen80-16     	     574	   2082874 ns/op	   50212 B/op	     444 allocs/op
BenchmarkNewVerifyLen/BenchmarkNewVerifyLen160-16    	     574	   2095455 ns/op	   50402 B/op	     444 allocs/op
BenchmarkNewVerifyLen/BenchmarkNewVerifyLen320-16    	     561	   2117191 ns/op	   50744 B/op	     443 allocs/op
BenchmarkVerify-16                                   	     402	   2977834 ns/op	   97272 B/op	     863 allocs/op
PASS
```









