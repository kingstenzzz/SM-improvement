
*sm2-improvement*

**Benchmark**
~~~~
goos: windows
goarch: amd64
pkg: github.com/kingstenzzz/sm2-improvement/sm2
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
>>优化sm2s算法
加密解密
BenchmarkLessThan32_NISTP256-16            10000            114534 ns/op            3639 B/op         70 allocs/op
BenchmarkLessThan32_P256SM2-16              8558            131622 ns/op            3639 B/op         70 allocs/op
BenchmarkMoreThan32_NISTP256-16            10000            118013 ns/op            5336 B/op         96 allocs/op
BenchmarkMoreThan32_P256SM2-16              9212            136440 ns/op            5336 B/op         96 allocs/op
签名验证
BenchmarkSM2_Sig-16                        18819             63395 ns/op             609 B/op         12 allocs/op

>>同济SM2算法
BenchmarkTjfoc_LessThan32_Enc-16             501           2365507 ns/op          150313 B/op       3115 allocs/op
BenchmarkTjfoc_MoreThan32_Enc-16             511           2347875 ns/op          153231 B/op       3138 allocs/op
签名验证
BenchmarkTjfoc_Sig-16                        780           1537682 ns/op           84155 B/op       1727 allocs/op


加密长度-运行时间 测试
goos: windows
goarch: amd64
pkg: github.com/kingstenzzz/sm2-improvement/sm2
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkDecryptCount
BenchmarkDecryptCount/len20
BenchmarkDecryptCount/len20-16         	    8905	    130858 ns/op	    3683 B/op	      70 allocs/op
BenchmarkDecryptCount/len40
BenchmarkDecryptCount/len40-16         	    9255	    132766 ns/op	    4308 B/op	      85 allocs/op
BenchmarkDecryptCount/len80
BenchmarkDecryptCount/len80-16         	    8594	    135499 ns/op	    5285 B/op	      98 allocs/op
BenchmarkDecryptCount/len160
BenchmarkDecryptCount/len160-16        	    8576	    140701 ns/op	    6582 B/op	     122 allocs/op
BenchmarkDecryptCount/len320
BenchmarkDecryptCount/len320-16        	    7519	    152538 ns/op	   10200 B/op	     188 allocs/op
BenchmarkDecryptCount/len640
BenchmarkDecryptCount/len640-16        	    6673	    176575 ns/op	   17084 B/op	     306 allocs/op
BenchmarkDecryptCount/len1280
BenchmarkDecryptCount/len1280-16       	    5461	    223094 ns/op	   30870 B/op	     538 allocs/op
BenchmarkDecryptCount/len2560
BenchmarkDecryptCount/len2560-16       	    3880	    317198 ns/op	   58794 B/op	    1018 allocs/op
BenchmarkDecryptCount/len5120
BenchmarkDecryptCount/len5120-16       	    2359	    505920 ns/op	  114896 B/op	    1978 allocs/op
BenchmarkDecryptCount/len10240
BenchmarkDecryptCount/len10240-16      	    1351	    874799 ns/op	  226199 B/op	    3898 allocs/op
>>>同济sm2
BenchmarkTjfocCount
BenchmarkTjfocCount/count52
BenchmarkTjfocCount/count52-16         	     762	   1544240 ns/op
BenchmarkTjfocCount/count208
BenchmarkTjfocCount/count208-16        	     781	   1548207 ns/op
BenchmarkTjfocCount/count832
BenchmarkTjfocCount/count832-16        	     775	   1546693 ns/op
BenchmarkTjfocCount/count3328
BenchmarkTjfocCount/count3328-16       	     763	   1580443 ns/op
BenchmarkTjfocCount/count13312
BenchmarkTjfocCount/count13312-16      	     704	   1707701 ns/op
BenchmarkTjfocCount/count53248
BenchmarkTjfocCount/count53248-16      	     522	   2209071 ns/op
BenchmarkTjfocCount/count212992
BenchmarkTjfocCount/count212992-16     	     283	   4231244 ns/op
BenchmarkTjfocCount/count851968
BenchmarkTjfocCount/count851968-16     	      92	  12177413 ns/op
BenchmarkTjfocCount/count3407872
BenchmarkTjfocCount/count3407872-16    	      26	  43425700 ns/op
BenchmarkTjfocCount/count13631488
BenchmarkTjfocCount/count13631488-16   	       6	 167843617 ns/op


Process finished with exit code 0


~~~~
## API使用示例


```Go
import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/kingstenzzz/sm2-improvement/sm2"
	"log"
)


func main()  {
	msg := []byte("test encryption")
	priv, err := sm2.GenerateKey(rand.Reader)
	if err!=nil {
		log.Fatal(err)

	}
	ciphertext, err := sm2.Encrypt(rand.Reader, &priv.PublicKey, msg, nil)
	fmt.Printf("加密结果:%x\n",ciphertext)
	plaintext,err := sm2.Decrypt(priv, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("解密结果",string(plaintext))
	if !bytes.Equal(msg,plaintext){
	log.Fatalf("解密失败")
	}

	r, s, err := sm2.SignWithSM2(rand.Reader, &priv.PrivateKey, nil,msg)
	if err != nil {
		log.Fatal("签名失败 %v", err)
	}
	result := sm2.VerifyWithSM2(&priv.PublicKey,nil, msg, r, s)
	if !result {
		log.Fatal("签名验证失败" )
	}
	fmt.Println("签名验证成功")
}
```

