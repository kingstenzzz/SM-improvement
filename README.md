
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

