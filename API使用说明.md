# 国密API使用说明
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
