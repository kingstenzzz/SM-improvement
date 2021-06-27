/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sm4

import (
	"fmt"
	"testing"
)

func TestSM4(t *testing.T) {
	key := []byte("1234567890abcdef")

	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	err := WriteKeyToPemFile("key.pem", key, nil)
	if err != nil {
		t.Fatalf("WriteKeyToPem error")
	}
	key, err = ReadKeyFromPemFile("key.pem", nil)
	fmt.Printf("key = %v\n", key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data = %x\n", data)
	ecbMsg, err := Sm4Ecb(key, data, true)
	if err != nil {
		t.Errorf("sm4 enc error:%s", err)
		return
	}
	fmt.Printf("ecbMsg = %x\n", ecbMsg)
	ecbDec, err := Sm4Ecb(key, ecbMsg, false)
	if err != nil {
		t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Printf("ecbDec = %x\n", ecbDec)
	if !testCompare(data, ecbDec) {
		t.Errorf("sm4 self enc and dec failed")
	}
	cbcMsg, err := Sm4Cbc(key, data, true)
	if err != nil {
		t.Errorf("sm4 enc error:%s", err)
	}
	fmt.Printf("cbcMsg = %x\n", cbcMsg)
	cbcDec, err := Sm4Cbc(key, cbcMsg, false)
	if err != nil {
		t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Printf("cbcDec = %x\n", cbcDec)
	if !testCompare(data, cbcDec) {
		t.Errorf("sm4 self enc and dec failed")
	}

	cbcMsg, err = Sm4CFB(key, data, true)
	if err != nil {
		t.Errorf("sm4 enc error:%s", err)
	}
	fmt.Printf("CFBMsg = %x\n", cbcMsg)

	cbcCfb, err := Sm4CFB(key, cbcMsg, false)
	if err != nil {
		t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Printf("CFBDec = %x\n", cbcCfb)

	cbcMsg, err = Sm4OFB(key, data, true)
	if err != nil {
		t.Errorf("sm4 enc error:%s", err)
	}
	fmt.Printf("OFBMsg = %x\n", cbcMsg)

	cbcOfc, err := Sm4OFB(key, cbcMsg, false)
	if err != nil {
		t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Printf("OFBDec = %x\n", cbcOfc)
}

/*func BenchmarkSM4(t *testing.B) {
	t.ReportAllocs()
	key := []byte("1234567890abcdef")
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	err := WriteKeyToPemFile("key.pem", key, nil)
	if err != nil {
		t.Fatalf("WriteKeyToPem error")
	}
	key, err = ReadKeyFromPemFile("key.pem", nil)
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < t.N; i++ {
		d0 := make([]byte, 16)
		c.Encrypt(d0, data)
		d1 := make([]byte, 16)
		c.Decrypt(d1, d0)
	}
}*/

func TestErrKeyLen(t *testing.T) {
	fmt.Printf("\n--------------test key len------------------")
	key := []byte("1234567890abcdefg")
	_, err := NewCipher(key)
	if err != nil {
		fmt.Println("\nError key len !")
	}
	key = []byte("1234")
	_, err = NewCipher(key)
	if err != nil {
		fmt.Println("Error key len !")
	}
	fmt.Println("------------------end----------------------")
}

func testCompare(key1, key2 []byte) bool {
	if len(key1) != len(key2) {
		return false
	}
	for i, v := range key1 {
		if i == 1 {
			//fmt.Println("type of v", reflect.TypeOf(v))
		}
		a := key2[i]
		if a != v {
			return false
		}
	}
	return true
}

func TestEnvior(t *testing.T) {
	key := []byte("1234567890abcdef")

	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	fmt.Printf("key = %v\n", key)
	fmt.Printf("data = %x\n", data)
	ecbMsg, err := Sm4Cbc(key, data, true) //Sm4Cbc模式pksc7填充加密
	if err != nil {
		//t.Errorf("sm4 enc error:%s", err)
		return
	}
	fmt.Printf("ecbMsg = %x\n", ecbMsg)
	ecbDec, err := Sm4Cbc(key, ecbMsg, false) //Sm4Cbc模式pksc7填充解密
	if err != nil {
		//t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Printf("ecbDec = %x\n", ecbDec)
	fmt.Printf("1111")
}
