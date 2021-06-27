package sm2

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"testing"

	"github.com/kingstenzzz/sm2-improvement/sm3"
	tjfoc_SM2 "github.com/tjfoc/gmsm/sm2" //引入同济库对比
)

func Test_kdf(t *testing.T) {
	x2, _ := new(big.Int).SetString("64D20D27D0632957F8028C1E024F6B02EDF23102A566C932AE8BD613A8E865FE", 16)
	y2, _ := new(big.Int).SetString("58D225ECA784AE300A81A2D48281A828E1CEDF11C4219099840265375077BF78", 16)

	expected := "006e30dae231b071dfad8aa379e90264491603"

	result, success := kdf(append(x2.Bytes(), y2.Bytes()...), 19)
	if !success {
		t.Fatalf("failed")
	}

	resultStr := hex.EncodeToString(result)

	if expected != resultStr {
		t.Fatalf("expected %s, real value %s", expected, resultStr)
	}
}

func Test_encryptDecrypt(t *testing.T) {
	priv, _ := GenerateKey(rand.Reader)
	tests := []struct {
		name      string
		plainText string
	}{
		// TODO: Add test cases.
		{"less than 32", "standardTS"},
		{"equals 32", "standardTS encryption "},
		{"long than 32", "standardTS standardTS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext, err := Encrypt(rand.Reader, &priv.PublicKey, []byte(tt.plainText), nil)
			if err != nil {
				t.Fatalf("encrypt failed %v", err)
			}
			plaintext, err := Decrypt(priv, ciphertext)
			if err != nil {
				t.Fatalf("decrypt failed %v", err)
			}
			if !reflect.DeepEqual(string(plaintext), tt.plainText) {
				t.Errorf("Decrypt() = %v, want %v", string(plaintext), tt.plainText)
			}
			// compress mode
			encrypterOpts := EncrypterOpts{MarshalCompressed}
			ciphertext, err = Encrypt(rand.Reader, &priv.PublicKey, []byte(tt.plainText), &encrypterOpts)
			fmt.Println("cpmpress mode")
			if err != nil {
				t.Fatalf("encrypt failed %v", err)
			}
			plaintext, err = Decrypt(priv, ciphertext)
			if err != nil {
				t.Fatalf("decrypt failed %v", err)
			}
			if !reflect.DeepEqual(string(plaintext), tt.plainText) {
				t.Errorf("Decrypt() = %v, want %v", string(plaintext), tt.plainText)
			}

			// mixed mode
			encrypterOpts = EncrypterOpts{MarshalMixed}
			fmt.Println("MarshalMixed mode")

			ciphertext, err = Encrypt(rand.Reader, &priv.PublicKey, []byte(tt.plainText), &encrypterOpts)
			if err != nil {
				t.Fatalf("encrypt failed %v", err)
			}
			plaintext, err = Decrypt(priv, ciphertext)
			if err != nil {
				t.Fatalf("decrypt failed %v", err)
			}
			if !reflect.DeepEqual(string(plaintext), tt.plainText) {
				t.Errorf("Decrypt() = %v, want %v", string(plaintext), tt.plainText)
			}
		})
	}
}

func Test_signVerify(t *testing.T) {
	priv, _ := GenerateKey(rand.Reader)
	tests := []struct {
		name      string
		plainText string
	}{
		// TODO: Add test cases.
		{"less than 32", "standardTS"},
		{"equals 32", "standardTS encryption "},
		{"long than 32", "standardTS standardTS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hash := sm3.Sm3Sum([]byte(tt.plainText))
			r, s, err := Sign(rand.Reader, &priv.PrivateKey, hash[:])
			if err != nil {
				t.Fatalf("sign failed %v", err)
			}
			result := Verify(&priv.PublicKey, hash[:], r, s)
			if !result {
				t.Fatal("verify failed")
			}
		})
	}
}

func benchmarkEncryptSM2(b *testing.B, plaintext string) {
	b.ReportAllocs()
	priv, _ := GenerateKey(rand.Reader)
	for i := 0; i < b.N; i++ {
		ciphertext, _ := Encrypt(rand.Reader, &priv.PublicKey, []byte(plaintext), nil)
		Decrypt(priv, ciphertext)
	}
}

func benchmarkEncryptNISTP256(b *testing.B, plaintext string) {
	b.ReportAllocs()
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ciphertext, _ := Encrypt(rand.Reader, &priv.PublicKey, []byte(plaintext), nil)
		Decrypt(&PrivateKey{*priv}, ciphertext)
	}
}

func BenchmarkLessThan32_NISTP256(b *testing.B) {
	benchmarkEncryptNISTP256(b, "standardTS")
}

func BenchmarkLessThan32_P256SM2(b *testing.B) {
	benchmarkEncryptSM2(b, "standardTS")
}

func BenchmarkMoreThan32_NISTP256(b *testing.B) {
	benchmarkEncryptNISTP256(b, "standardTS standardTS standardTS standardTS standardTS standardTS")
}

func BenchmarkMoreThan32_P256SM2(b *testing.B) {

	benchmarkEncryptSM2(b, "standard teststandard teststandard teststandard test")

}

func BenchmarkTjfoc_LessThan32_Enc(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("standardTS")
	priv, _ := tjfoc_SM2.GenerateKey(nil) // 生成密钥对
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		ciphertext, _ := priv.PublicKey.EncryptAsn1(msg, rand.Reader)
		priv.DecryptAsn1(ciphertext)
	}
}

func BenchmarkTjfoc_MoreThan32_Enc(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("standardTS standardTS standardTS standardTS standardTS standardTS")
	priv, _ := tjfoc_SM2.GenerateKey(nil) // 生成密钥对
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		ciphertext, _ := priv.PublicKey.EncryptAsn1(msg, rand.Reader)
		priv.DecryptAsn1(ciphertext)
	}
}

func BenchmarkDecryptCount(b *testing.B) {
	msg := "standardTS"
	for i := 0; i < 10; i++ {
		msg = msg + msg + msg

		b.Run("len"+strconv.Itoa(len(msg)), func(b *testing.B) {
			benchmarkEncryptSM2(b, msg)
		})
	}
}

func BenchmarkSM2_Sig(t *testing.B) {
	t.ReportAllocs()
	priv, _ := GenerateKey(rand.Reader)
	msg := []byte("standardTS")

	hash := sm3.Sm3Sum(msg)
	r, s, _ := Sign(rand.Reader, &priv.PrivateKey, hash[:])
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		Verify(&priv.PublicKey, hash[:], r, s) // 密钥验证
	}
}

func BenchmarkTjfoc_Sig(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("standardTS")
	priv, _ := tjfoc_SM2.GenerateKey(nil) // 生成密钥对
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sign, _ := priv.Sign(nil, msg, nil) // 签名
		priv.Verify(msg, sign)              // 密钥验证
	}
}
func BenchmarkTjfocCount(b *testing.B) {
	msg := "standardTS"
	priv, _ := tjfoc_SM2.GenerateKey(nil) // 生成密钥对
	for i := 0; i < 10; i++ {
		msg = msg + msg
		b.Run("len"+strconv.Itoa(len(msg)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ciphertext, _ := priv.PublicKey.EncryptAsn1([]byte(msg), rand.Reader)
				priv.DecryptAsn1(ciphertext)
			}
		})
	}
}
