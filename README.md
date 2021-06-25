
*sm2-improvement*

**Benchmark**
~~~~
goos: windows
goarch: amd64
pkg: github.com/kingstenzzz/sm2-improvement/sm2
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkSM2-16                                    18955             63025 ns/op             608 B/op         12 allocs/op
BenchmarkLessThan32_P256-16                        16273             73702 ns/op            2434 B/op         48 allocs/op
BenchmarkLessThan32_P256SM2-16                     13422             89487 ns/op            2434 B/op         48 allocs/op
BenchmarkMoreThan32_P256-16                        15246             78880 ns/op            4107 B/op         75 allocs/op
BenchmarkMoreThan32_P256SM2-16                     12694             94581 ns/op            4107 B/op         75 allocs/op
>>>>同济库
BenchmarkTjfoc_MoreThan32_P256SM2-16                 774           1547875 ns/op           84253 B/op       1725 allocs/op
BenchmarkTjfoc_LessThan32_P256SM2-16                 790           1528133 ns/op           83975 B/op       1724 allocs/op
PASS

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

