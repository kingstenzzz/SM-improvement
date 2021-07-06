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
[README.md](https://github.com/LYL20200307/SM-improvement/files/6767097/README.md)
