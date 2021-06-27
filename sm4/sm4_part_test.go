package sm4

import (
	tjfoc "github.com/tjfoc/gmsm/sm4"
	"testing"
)

/*func BenchmarkPart1(t *testing.B){
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		key := []byte("1234567890abcdef")
		b := make([]uint32, 4)
		b[0] = binary.BigEndian.Uint32(key[0:4])
		b[1] = binary.BigEndian.Uint32(key[4:8])
		b[2] = binary.BigEndian.Uint32(key[8:12])
		b[3] = binary.BigEndian.Uint32(key[12:16])
	}


}


func BenchmarkPart2(t *testing.B){
	//big:=NewBig()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		block := []byte("1234567890abcdef")
		b := make([]uint32, 4)
		for i := 0; i < 4; i++ {
			b[i] = (uint32(block[i*4]) << 24) | (uint32(block[i*4+1]) << 16) |
				(uint32(block[i*4+2]) << 8) | (uint32(block[i*4+3]))
		}
	}



}*/
//
//
//func TestPart1(t *testing.T){
//	block := []byte("1234567890abcdef")
//	b := make([]uint32, 4)
//	for i := 0; i < 4; i++ {
//		b[i] = (uint32(block[i*4]) << 24) | (uint32(block[i*4+1]) << 16) |
//			(uint32(block[i*4+2]) << 8) | (uint32(block[i*4+3]))
//	}
//	fmt.Printf("data = %b\n", uint32(block[0])<< 24)
//	fmt.Printf("data = %b\n", uint32(block[1])<< 16)
//	fmt.Printf("data = %b\n", uint32(block[2])<< 8)
//	fmt.Printf("data = %b\n", uint32(block[3]))
//	fmt.Printf("data = %b\n", b[0])
//	fmt.Printf("data = %d\n", b[1])
//	fmt.Printf("data = %d\n", b[2])
//	fmt.Printf("data = %d\n", b[3])
//
//}
//func TestPart2(t *testing.T){
//	key := []byte("1234567890abcdef")
//	b := make([]uint32, 4)
//	b[0] = binary.BigEndian.Uint32(key[0:4])
//	b[1] = binary.BigEndian.Uint32(key[4:8])
//	b[2] = binary.BigEndian.Uint32(key[8:12])
//	b[3] = binary.BigEndian.Uint32(key[12:16])
//	fmt.Printf("data = %d\n", b[0])
//	fmt.Printf("data = %d\n", b[1])
//	fmt.Printf("data = %d\n", b[2])
//	fmt.Printf("data = %d\n", b[3])
//
//
//}
//
//
//

/*func BenchmarkPart3(t *testing.B){
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")
		subkeys := make([]uint32, 32)
		b := make([]uint32, 4)
		permuteInitialBlock1(b, key)
		b[0] ^= fk[0]
		b[1] ^= fk[1]
		b[2] ^= fk[2]
		b[3] ^= fk[3]
		for i := 0; i < 32; i++ {
			subkeys[i] = feistel1(b[0], b[1], b[2], b[3], ck[i])
			b[0], b[1], b[2], b[3] = b[1], b[2], b[3], subkeys[i]
		}
	}




}

func BenchmarkPart4(t *testing.B) {
	//big:=NewBig()
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")
		subkeys := make([]uint32, 32)
		b := make([]uint32, 4)
		permuteInitialBlock(b, key)
		b[0] ^= fk[0]
		b[1] ^= fk[1]
		b[2] ^= fk[2]
		b[3] ^= fk[3]
		for i := 0; i < 32; i++ {
			subkeys[i] = feistel0(b[0], b[1], b[2], b[3], ck[i])
			b[0], b[1], b[2], b[3] = b[1], b[2], b[3], subkeys[i]
		}
	}
}*/
//
//
//func TestPart3(t *testing.T){
//		key := []byte("1234567890abcdef")
//		subkeys := make([]uint32, 32)
//		b := make([]uint32, 4)
//		permuteInitialBlock(b, key)
//		b[0] ^= fk[0]
//		b[1] ^= fk[1]
//		b[2] ^= fk[2]
//		b[3] ^= fk[3]
//		for i := 0; i < 32; i++ {
//			subkeys[i] = feistel1(b[0], b[1], b[2], b[3], ck[i])
//			b[0], b[1], b[2], b[3] = b[1], b[2], b[3], subkeys[i]
//		}
//		fmt.Printf("subkeys = %d\n", subkeys)
//}
//
//
//func TestPart4(t *testing.T) {
//		key := []byte("1234567890abcdef")
//		subkeys := make([]uint32, 32)
//		b := make([]uint32, 4)
//		permuteInitialBlock(b, key)
//		b[0] ^= fk[0]
//		b[1] ^= fk[1]
//		b[2] ^= fk[2]
//		b[3] ^= fk[3]
//		for i := 0; i < 32; i++ {
//			subkeys[i] = feistel0(b[0], b[1], b[2], b[3], ck[i])
//			b[0], b[1], b[2], b[3] = b[1], b[2], b[3], subkeys[i]
//		}
//		fmt.Printf("subkeys = %d\n", subkeys)
//}

/*func BenchmarkPart5_1(t *testing.B){
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		b := make([]uint32, 4)
		block:=[]byte{49,50,51 ,52 ,53 ,54, 55, 56, 57 ,48, 97 ,98 ,99, 100, 101, 102}
		b[0] = ToUint32(block[0:4])
		b[1] = ToUint32(block[4:8])
		b[2] = ToUint32(block[8:12])
		b[3] = ToUint32(block[12:16])
	}
}
func BenchmarkPart5_2(t *testing.B){
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		b := make([]uint32, 4)
		block:=[]byte{49,50,51 ,52 ,53 ,54, 55, 56, 57 ,48, 97 ,98 ,99, 100, 101, 102}
		b[0] = binary.BigEndian.Uint32(block[0:4])
		b[1] = binary.BigEndian.Uint32(block[4:8])
		b[2] = binary.BigEndian.Uint32(block[8:12])
		b[3] = binary.BigEndian.Uint32(block[12:16])
	}
}




func BenchmarkPart6(t *testing.B) {
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		b := make([]uint32, 4)
		block:=[]byte{49,50,51 ,52 ,53 ,54, 55, 56, 57 ,48, 97 ,98 ,99, 100, 101, 102}
		for i := 0; i < 4; i++ {
			b[i] = (uint32(block[i*4]) << 24) | (uint32(block[i*4+1]) << 16) |
				(uint32(block[i*4+2]) << 8) | (uint32(block[i*4+3]))
		}
	}

}
func BenchmarkPart7(t *testing.B){
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		b := []byte{49,50,51 ,52 ,53 ,54, 55, 56, 57 ,48, 97 ,98 ,99, 100, 101, 102}
		block:=[]uint32{19088743,2309737967,4275878552,1985229328}
		binary.BigEndian.PutUint32(b[0:4], block[0])
		binary.BigEndian.PutUint32(b[4:8], block[1])
		binary.BigEndian.PutUint32(b[8:12], block[2])
		binary.BigEndian.PutUint32(b[12:16], block[3])
	}
}

func BenchmarkPart8(t *testing.B) {
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		b := []byte{49,50,51 ,52 ,53 ,54, 55, 56, 57 ,48, 97 ,98 ,99, 100, 101, 102}
		block:=[]uint32{19088743,2309737967,4275878552,1985229328}
		for i := 0; i < 4; i++ {
			b[i*4] = uint8(block[i] >> 24)
			b[i*4+1] = uint8(block[i] >> 16)
			b[i*4+2] = uint8(block[i] >> 8)
			b[i*4+3] = uint8(block[i])
		}
	}

}
func BenchmarkPart(t *testing.B) {
	t.ReportAllocs()
	for j := 0; j < t.N; j++ {
		key := []byte("1234567890abcdef")
		//fmt.Printf("key = %v\n", key)

		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
		//fmt.Printf("data = %x\n", data)
		ecbMsg, err :=Sm4Cbc(key, data, true)   //Sm4Cbc模式pksc7填充加密
		if err != nil {
			//t.Errorf("sm4 enc error:%s", err)
			return
		}
		//fmt.Printf("ecbMsg = %x\n", ecbMsg)
		ecbDec, err := Sm4Cbc(key, ecbMsg, false)  //Sm4Cbc模式pksc7填充解密
		if err != nil {
			//t.Errorf("sm4 dec error:%s", err)
			return
		}
		if ecbDec == nil {
			//t.Errorf("sm4 dec error:%s", err)
			return
		}


		//fmt.Printf("ecbDec = %x\n", ecbDec)

	}

}*/
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
