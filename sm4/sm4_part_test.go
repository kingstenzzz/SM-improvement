package sm4

import (
	tjfoc "github.com/tjfoc/gmsm/sm4"
	"testing"
)

func BenchmarkPart_Ecb_tj(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		ecbMsg, err := tjfoc.Sm4Ecb(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
			return
		}
		ecbDec, err := tjfoc.Sm4Ecb(key, ecbMsg, false)
		if err != nil || ecbDec == nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}

	}
}
func BenchmarkPart_Cbc_tj(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		cbcMsg, err := tjfoc.Sm4Cbc(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
		}
		cbcDec, err := tjfoc.Sm4Cbc(key, cbcMsg, false)
		if err != nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}
		if !testCompare(data, cbcDec) {
			t.Errorf("sm4 self enc and dec failed")
		}
	}
}

func BenchmarkPart_Cfb_tj(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		cfbMsg, err := tjfoc.Sm4CFB(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
		}

		cbcCfb, err := tjfoc.Sm4CFB(key, cfbMsg, false)
		if err != nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}
		if !testCompare(data, cbcCfb) {
			t.Errorf("sm4 self enc and dec failed")
		}

	}
}

func BenchmarkPart_Ofb_tj(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		OfbMsg, err := tjfoc.Sm4OFB(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
		}

		cbcOfc, err := tjfoc.Sm4OFB(key, OfbMsg, false)
		if err != nil || cbcOfc == nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}
		if !testCompare(data, cbcOfc) {
			t.Errorf("sm4 self enc and dec failed")
		}

	}
}

func BenchmarkPart_Ecb(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		ecbMsg, err := Sm4Ecb(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
			return
		}
		ecbDec, err := Sm4Ecb(key, ecbMsg, false)
		if err != nil || ecbDec == nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}

	}
}

func BenchmarkPart_Cbc(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		cbcMsg, err := Sm4Cbc(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
		}
		cbcDec, err := Sm4Cbc(key, cbcMsg, false)
		if err != nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}
		if !testCompare(data, cbcDec) {
			t.Errorf("sm4 self enc and dec failed")
		}
	}
}

func BenchmarkPart_Cfb(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		cfbMsg, err := Sm4CFB(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
		}

		cbcCfb, err := Sm4CFB(key, cfbMsg, false)
		if err != nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}
		if !testCompare(data, cbcCfb) {
			t.Errorf("sm4 self enc and dec failed")
		}

	}
}

func BenchmarkPart_Ofb(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		OfbMsg, err := Sm4OFB(key, data, true)
		if err != nil {
			t.Errorf("sm4 enc error:%s", err)
		}

		cbcOfc, err := Sm4OFB(key, OfbMsg, false)
		if err != nil || cbcOfc == nil {
			t.Errorf("sm4 dec error:%s", err)
			return
		}
		if !testCompare(data, cbcOfc) {
			t.Errorf("sm4 self enc and dec failed")
		}

	}
}
