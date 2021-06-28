
# sm-improvement

~~~
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
SM2：
Benchmark_P256SM2-16              9212            136440 ns/op            5336 B/op         96 allocs/op
签名验证
BenchmarkSM2_Sig-16               18819            63395 ns/op             609 B/op         12 allocs/op

SM3：
BenchmarkSm3-16                 18065758              66.40 ns/op           19 B/op          0 allocs/op

SM4:
BenchmarkPart_Ecb-16              960276              1243 ns/op             672 B/op         17 allocs/op
BenchmarkPart_Cbc-16              925554              1338 ns/op             736 B/op         21 allocs/op
BenchmarkPart_Cfb-16              925525              1345 ns/op             768 B/op         23 allocs/op
BenchmarkPart_Ofb-16              923403              1337 ns/op             752 B/op         22 allocs/op

SM9:
BenchmarkNewSignVerify-16    	     393	   3017218 ns/op	   52656 B/op	     467 allocs/op

~~~
