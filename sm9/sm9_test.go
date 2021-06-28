// Copyright 2020 cetc-30. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm9

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	mk, err := MasterKeyGen(rand.Reader)
	if err != nil {
		t.Errorf("mk gen failed:%s", err)
		return
	}
	var hid byte = 1
	var uid = []byte("Alice")

	uk, err := UserKeyGen(mk, uid, hid)
	if err != nil {
		t.Errorf("uk gen failed:%s", err)
		return
	}

	msg := []byte("message")

	sig, err := Sign(uk, &mk.MasterPubKey, msg)
	if err != nil {
		t.Errorf("sm9 sign failed:%s", err)
		return
	}
	if !Verify(sig, msg, uid, hid, &mk.MasterPubKey) {
		t.Error("sm9 sig is invalid")
		return
	}
}

func TestNewSignAndVerify(t *testing.T) {
	mk, err := MasterKeyGen(rand.Reader)
	if err != nil {
		t.Errorf("mk gen failed:%s", err)
		return
	}
	var hid byte = 1
	var uid = []byte("Alice")

	uk, err := UserKeyGen(mk, uid, hid)
	if err != nil {
		t.Errorf("uk gen failed:%s", err)
		return
	}

	msg := []byte("message")

	sig, err := NewSign(uk, &mk.MasterPubKey, msg)
	if err != nil {
		t.Errorf("sm9 sign failed:%s", err)
		return
	}
	if !Verify(sig, msg, uid, hid, &mk.MasterPubKey) {
		t.Error("sm9 sig is invalid")
		return
	}
}

func TestSignAndNewVerify(t *testing.T) {
	mk, err := MasterKeyGen(rand.Reader)
	if err != nil {
		t.Errorf("mk gen failed:%s", err)
		return
	}
	var hid byte = 1
	var uid = []byte("Alice")

	uk, err := UserKeyGen(mk, uid, hid)
	if err != nil {
		t.Errorf("uk gen failed:%s", err)
		return
	}

	msg := []byte("message")

	sig, err := Sign(uk, &mk.MasterPubKey, msg)
	if err != nil {
		t.Errorf("sm9 sign failed:%s", err)
		return
	}
	if !NewVerify(sig, msg, uid, hid, &mk.MasterPubKey) {
		t.Error("sm9 sig is invalid")
		return
	}
}

func TestNewSignAndNewVerify(t *testing.T) {
	mk, err := MasterKeyGen(rand.Reader)
	if err != nil {
		t.Errorf("mk gen failed:%s", err)
		return
	}
	var hid byte = 1
	var uid = []byte("Alice")

	uk, err := UserKeyGen(mk, uid, hid)
	if err != nil {
		t.Errorf("uk gen failed:%s", err)
		return
	}

	msg := []byte("message")

	sig, err := NewSign(uk, &mk.MasterPubKey, msg)
	if err != nil {
		t.Errorf("sm9 sign failed:%s", err)
		return
	}
	if !NewVerify(sig, msg, uid, hid, &mk.MasterPubKey) {
		t.Error("sm9 sig is invalid")
		return
	}
}

func BenchmarkSign(b *testing.B) {
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Sign(uk, &mk.MasterPubKey, msg)
	}
}

func BenchmarkNewSign(b *testing.B) {
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewSign(uk, &mk.MasterPubKey, msg)
	}
}

func BenchmarkNewVerify(b *testing.B) {
	var State bool
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message")

	sig, _ := Sign(uk, &mk.MasterPubKey, msg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		State = NewVerify(sig, msg, id, byte(hid), &mk.MasterPubKey)
		if State != true {
			b.Errorf("sm9 sig is invalid")
		}
	}
	b.StopTimer()
}

func BenchmarkNewSignVerify(b *testing.B) {
	var State bool
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sig, _ := NewSign(uk, &mk.MasterPubKey, msg)
		State = NewVerify(sig, msg, id, byte(hid), &mk.MasterPubKey)
		if State != true {
			b.Errorf("sm9 sig is invalid")
		}
	}
	b.StopTimer()
}

func BenchmarkNewSignLen(b *testing.B) {
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message111")

	for i := 0; i < 5; i++ {
		msg = append(msg, msg...)
		name := fmt.Sprintf("BenchmarkNewSignLen%v", len(msg))
		b.Run(name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				_, _ = NewSign(uk, &mk.MasterPubKey, msg)
			}
		})
	}
}

func BenchmarkNewVerifyLen(b *testing.B) {
	var State bool
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message111")

	for i := 0; i < 5; i++ {
		msg = append(msg, msg...)
		sig, _ := Sign(uk, &mk.MasterPubKey, msg)
		name := fmt.Sprintf("BenchmarkNewVerifyLen%v", len(msg))
		b.Run(name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				State = NewVerify(sig, msg, id, byte(hid), &mk.MasterPubKey)
				if State != true {
					b.Errorf("sm9 sig is invalid")
				}
			}
		})
	}

}

func BenchmarkVerify(b *testing.B) {
	var State bool
	mk, _ := MasterKeyGen(rand.Reader)
	id := []byte("Alice")
	hid := 3
	uk, _ := UserKeyGen(mk, id, byte(hid))

	var msg = []byte("message")

	sig, _ := Sign(uk, &mk.MasterPubKey, msg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		State = Verify(sig, msg, id, byte(hid), &mk.MasterPubKey)
		if State != true {
			b.Errorf("sm9 sig is invalid")
		}
	}
	b.StopTimer()
}
