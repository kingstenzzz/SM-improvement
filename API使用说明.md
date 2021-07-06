# 国密Go API使用说明
## Go包安装

```bash
go get -u github.com/kingstenzzz/SM-improvement
```
### SM2示例
- go package： `github.com/kingstenzzz/SM-improvement/sm2`
```Go
	msg := []byte("test encryption")
	priv, err := sm2.GenerateKey(rand.Reader)
	if err!=nil {
		log.Fatal(err)

	}
	ciphertext, err := sm2.Encrypt(rand.Reader, &priv.PublicKey, msg, nil)//加密
	fmt.Printf("加密结果:%x\n",ciphertext)
	plaintext,err := sm2.Decrypt(priv, ciphertext)//解密
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("解密结果",string(plaintext))
	if !bytes.Equal(msg,plaintext){
	log.Fatalf("解密失败")
	}

	r, s, err := sm2.SignWithSM2(rand.Reader, &priv.PrivateKey, nil,msg)//签名
	if err != nil {
		log.Fatal("签名失败 %v", err)
	}
	result := sm2.VerifyWithSM2(&priv.PublicKey,nil, msg, r, s)//验证
	if !result {
		log.Fatal("签名验证失败" )
	}
	fmt.Println("签名验证成功")
```
## SM3密码杂凑算法示例

```Go
    data := "test"
    h := sm3.New()
    h.Write([]byte(data))
    sum := h.Sum(nil)
    fmt.Printf("digest value is: %x\n",sum)
```

## SM4分组密码算法 - 

- go package：`github.com/kingstenzzz/SM-improvement/sm4`

### 代码示例

```Go
    import  "crypto/cipher"
    import  "github.com/kingstenzzz/SM-improvement/sm4"
    import "fmt"

    func main(){
    key := []byte("1234567890abcdef")
	fmt.Printf("key = %v\n", key)
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
    iv := []byte("0000000000000000")
	err = SetIV(iv)//设置SM4算法实现的IV值,不设置则使用默认值
	ecbMsg, err :=sm4.Sm4Ecb(key, data, true)   //sm4Ecb模式pksc7填充加密
	if err != nil {
		t.Errorf("sm4 enc error:%s", err)
		return
	}
	fmt.Printf("ecbMsg = %x\n", ecbMsg)
	ecbDec, err := sm4.Sm4Ecb(key, ecbMsg, false)  //sm4Ecb模式pksc7填充解密
	if err != nil {
		t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Printf("ecbDec = %x\n", ecbDec)
    }
```




### SM9 示例
```go
	mk, err := sm9.MasterKeyGen(rand.Reader) //生成系统签名主密钥对
	if err != nil {
		log.Fatalf("mk gen failed:%s", err)
	}
	var hid byte = 1//签名私钥生成函数识别符
	var uid = []byte("Alice")//用户标识

	uk, err := sm9.UserKeyGen(mk, uid, hid)//生成用户签名私钥
	if err != nil {
		log.Fatalf("uk gen failed:%s", err)
	}

	msg := []byte("message")

	sig, err := sm9.NewSign(uk, &mk.MasterPubKey, msg)
	if err != nil {
		log.Fatalf("sm9 签名失败:%s", err)
	}

	result := sm9.NewVerify(sig, msg, uid, hid, &mk.MasterPubKey)
	if !result {
		log.Fatal("签名验证失败")
	}
	fmt.Println("签名验证成功")
```


### 具体功能测试代码参考
```Go
github.com/kingstenzzz/SM-improvement/sm2/sm2_test.go  //sm3算法
github.com/kingstenzzz/SM-improvement/sm3/sm3_test.go  //sm3算法
github.com/kingstenzzz/SM-improvement/sm4/sm4_test.go  //sm3算法
github.com/kingstenzzz/SM-improvement/sm9/sm9_test.go  //sm3算法
```
