package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/kingstenzzz/sm2-improvement/sm2"
	"github.com/kingstenzzz/sm2-improvement/sm3"

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

	hash := sm3.Sum(msg)
	r, s, err := sm2.Sign(rand.Reader, &priv.PrivateKey, hash[:])
	if err != nil {
		log.Fatal("签名失败 %v", err)
	}
	result := sm2.Verify(&priv.PublicKey, hash[:], r, s)
	if !result {
		log.Fatal("签名验证失败" )
	}
	fmt.Println("签名验证成功")
}