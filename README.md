
*sm2-improvement*

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

