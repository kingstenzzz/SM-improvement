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

package sm3

import (
	"encoding/binary"
	"fmt"
	"hash"
	"time"
)

// ti为压缩函数64轮中每一轮 Ti<<<i 的结果
var ti = [64]uint32{
	2043430169, 4086860338, 3878753381, 3462539467, 2630111639, 965255983, 1930511966, 3861023932,
	3427080569, 2559193843, 823420391, 1646840782, 3293681564, 2292395833, 289824371, 579648742,
	2643098247, 991229199, 1982458398, 3964916796, 3634866297, 2974765299, 1654563303, 3309126606,
	2323285917, 351604539, 703209078, 1406418156, 2812836312, 1330705329, 2661410658, 1027854021,
	2055708042, 4111416084, 3927864873, 3560762451, 2826557607, 1358147919, 2716295838, 1137624381,
	2275248762, 255530229, 511060458, 1022120916, 2044241832, 4088483664, 3882000033, 3469032771,
	2643098247, 991229199, 1982458398, 3964916796, 3634866297, 2974765299, 1654563303, 3309126606,
	2323285917, 351604539, 703209078, 1406418156, 2812836312, 1330705329, 2661410658, 1027854021}

type SM3 struct {
	digest      [8]uint32 // digest represents the partial evaluation of V
	length      uint64    // length of the message
	unhandleMsg []byte    // uint8  //
}

func (sm3 *SM3) ff0(x, y, z uint32) uint32 { return x ^ y ^ z }

func (sm3 *SM3) ff1(x, y, z uint32) uint32 { return (x & y) | (x & z) | (y & z) }

func (sm3 *SM3) gg0(x, y, z uint32) uint32 { return x ^ y ^ z }

func (sm3 *SM3) gg1(x, y, z uint32) uint32 { return (x & y) | (^x & z) }

func (sm3 *SM3) p0(x uint32) uint32 { return x ^ sm3.leftRotate(x, 9) ^ sm3.leftRotate(x, 17) }

func (sm3 *SM3) p1(x uint32) uint32 { return x ^ sm3.leftRotate(x, 15) ^ sm3.leftRotate(x, 23) }

func (sm3 *SM3) leftRotate(x uint32, i uint32) uint32 { return x<<(i%32) | x>>(32-i%32) }

// 填充
func (sm3 *SM3) pad() []byte {
	msg := sm3.unhandleMsg
	msg = append(msg, 0x80) // Append '1'
	blockSize := 64         // Append until the resulting message length (in bits) is congruent to 448 (mod 512)
	for len(msg)%blockSize != 56 {
		msg = append(msg, 0x00)
	}
	// append message length
	msg = append(msg, uint8(sm3.length>>56&0xff))
	msg = append(msg, uint8(sm3.length>>48&0xff))
	msg = append(msg, uint8(sm3.length>>40&0xff))
	msg = append(msg, uint8(sm3.length>>32&0xff))
	msg = append(msg, uint8(sm3.length>>24&0xff))
	msg = append(msg, uint8(sm3.length>>16&0xff))
	msg = append(msg, uint8(sm3.length>>8&0xff))
	msg = append(msg, uint8(sm3.length>>0&0xff))

	if len(msg)%64 != 0 {
		panic("------SM3 Pad: error msgLen =")
	}
	return msg
}

func (sm3 *SM3) update3(msg []byte) {
	var w [68]uint32
	var i int
	var TT1, TT2 uint32
	var A, B, C, D, E, F, G, H uint32

	a, b, c, d, e, f, g, h := sm3.digest[0], sm3.digest[1], sm3.digest[2], sm3.digest[3], sm3.digest[4], sm3.digest[5], sm3.digest[6], sm3.digest[7]
	for len(msg) >= 64 {
		for i = 0; i < 4; i++ {
			w[i] = binary.BigEndian.Uint32(msg[4*i : 4*(i+1)])
		}
		A, B, C, D, E, F, G, H = a, b, c, d, e, f, g, h
		for i = 0; i < 12; i++ {
			w[i+4] = binary.BigEndian.Uint32(msg[4*(i+4) : 4*(i+5)])

			TT1 = sm3.ff0(A, B, C) + D + ((sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7)) ^ sm3.leftRotate(A, 12)) + (w[i] ^ w[i+4])
			TT2 = sm3.gg0(E, F, G) + H + sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7) + w[i]
			D = C
			C = sm3.leftRotate(B, 9)
			B = A
			A = TT1
			H = G
			G = sm3.leftRotate(F, 19)
			F = E
			E = sm3.p0(TT2)
		}
		for i = 12; i < 16; i++ {
			w[i+4] = sm3.p1(w[i-12]^w[i-5]^sm3.leftRotate(w[i+1], 15)) ^ sm3.leftRotate(w[i-9], 7) ^ w[i-2]

			TT1 = sm3.ff0(A, B, C) + D + ((sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7)) ^ sm3.leftRotate(A, 12)) + (w[i] ^ w[i+4])
			TT2 = sm3.gg0(E, F, G) + H + sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7) + w[i]
			D = C
			C = sm3.leftRotate(B, 9)
			B = A
			A = TT1
			H = G
			G = sm3.leftRotate(F, 19)
			F = E
			E = sm3.p0(TT2)
		}

		for i := 16; i < 64; i++ {
			w[i+4] = sm3.p1(w[i-12]^w[i-5]^sm3.leftRotate(w[i+1], 15)) ^ sm3.leftRotate(w[i-9], 7) ^ w[i-2]

			TT1 = sm3.ff1(A, B, C) + D + ((sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7)) ^ sm3.leftRotate(A, 12)) + (w[i] ^ w[i+4])
			TT2 = sm3.gg1(E, F, G) + H + sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7) + w[i]

			D = C
			C = sm3.leftRotate(B, 9)
			B = A
			A = TT1
			H = G
			G = sm3.leftRotate(F, 19)
			F = E
			E = sm3.p0(TT2)
		}
		a ^= A
		b ^= B
		c ^= C
		d ^= D
		e ^= E
		f ^= F
		g ^= G
		h ^= H
		msg = msg[64:]
	}
	sm3.digest[0], sm3.digest[1], sm3.digest[2], sm3.digest[3], sm3.digest[4], sm3.digest[5], sm3.digest[6], sm3.digest[7] = a, b, c, d, e, f, g, h
}

func (sm3 *SM3) update4(msg []byte) [8]uint32 {
	var w [68]uint32
	var i int
	var TT1, TT2 uint32
	var A, B, C, D, E, F, G, H uint32

	a, b, c, d, e, f, g, h := sm3.digest[0], sm3.digest[1], sm3.digest[2], sm3.digest[3], sm3.digest[4], sm3.digest[5], sm3.digest[6], sm3.digest[7]
	for len(msg) >= 64 {
		for i = 0; i < 4; i++ {
			w[i] = binary.BigEndian.Uint32(msg[4*i : 4*(i+1)])
		}
		A, B, C, D, E, F, G, H = a, b, c, d, e, f, g, h
		for i = 0; i < 12; i++ {
			w[i+4] = binary.BigEndian.Uint32(msg[4*(i+4) : 4*(i+5)])

			TT1 = sm3.ff0(A, B, C) + D + ((sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7)) ^ sm3.leftRotate(A, 12)) + (w[i] ^ w[i+4])
			TT2 = sm3.gg0(E, F, G) + H + sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7) + w[i]
			D = C
			C = sm3.leftRotate(B, 9)
			B = A
			A = TT1
			H = G
			G = sm3.leftRotate(F, 19)
			F = E
			E = sm3.p0(TT2)
		}
		for i = 12; i < 16; i++ {
			w[i+4] = sm3.p1(w[i-12]^w[i-5]^sm3.leftRotate(w[i+1], 15)) ^ sm3.leftRotate(w[i-9], 7) ^ w[i-2]

			TT1 = sm3.ff0(A, B, C) + D + ((sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7)) ^ sm3.leftRotate(A, 12)) + (w[i] ^ w[i+4])
			TT2 = sm3.gg0(E, F, G) + H + sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7) + w[i]
			D = C
			C = sm3.leftRotate(B, 9)
			B = A
			A = TT1
			H = G
			G = sm3.leftRotate(F, 19)
			F = E
			E = sm3.p0(TT2)
		}

		for i := 16; i < 64; i++ {
			w[i+4] = sm3.p1(w[i-12]^w[i-5]^sm3.leftRotate(w[i+1], 15)) ^ sm3.leftRotate(w[i-9], 7) ^ w[i-2]

			TT1 = sm3.ff1(A, B, C) + D + ((sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7)) ^ sm3.leftRotate(A, 12)) + (w[i] ^ w[i+4])
			TT2 = sm3.gg1(E, F, G) + H + sm3.leftRotate(sm3.leftRotate(A, 12)+E+ti[i], 7) + w[i]

			D = C
			C = sm3.leftRotate(B, 9)
			B = A
			A = TT1
			H = G
			G = sm3.leftRotate(F, 19)
			F = E
			E = sm3.p0(TT2)
		}
		a ^= A
		b ^= B
		c ^= C
		d ^= D
		e ^= E
		f ^= F
		g ^= G
		h ^= H
		msg = msg[64:]
	}
	var digest [8]uint32
	digest[0], digest[1], digest[2], digest[3], digest[4], digest[5], digest[6], digest[7] = a, b, c, d, e, f, g, h
	return digest
}

// 创建哈希计算实例
func New() hash.Hash {
	var sm3 SM3

	sm3.Reset()
	return &sm3
}

// BlockSize 返回哈希的底层块大小。
func (sm3 *SM3) BlockSize() int { return 64 }

// Size returns the number of bytes Sum will return.
// Size 返回 Sum 将返回的字节数。
func (sm3 *SM3) Size() int { return 32 }

// Reset clears the internal state by zeroing bytes in the state buffer.
// This can be skipped for a newly-created hash state; the default zero-allocated state is correct.
// 重置将哈希重置为其初始状态。
func (sm3 *SM3) Reset() {
	// Reset digest
	sm3.digest[0] = 0x7380166f
	sm3.digest[1] = 0x4914b2b9
	sm3.digest[2] = 0x172442d7
	sm3.digest[3] = 0xda8a0600
	sm3.digest[4] = 0xa96f30bc
	sm3.digest[5] = 0x163138aa
	sm3.digest[6] = 0xe38dee4d
	sm3.digest[7] = 0xb0fb0e4e

	sm3.length = 0 // Reset numberic states
	sm3.unhandleMsg = []byte{}
}

// 嵌入了 io.Writer 接口
// 向执行中的 hash 加入更多数据
// hash 函数的 Write 方法永远不会返回 error
// Write方法必须能够接受任何数量的数据，但如果所有写入是块大小的倍数，它可以更有效地运行。
func (sm3 *SM3) Write(p []byte) (int, error) {
	toWrite := len(p)
	sm3.length += uint64(len(p) * 8)
	msg := append(sm3.unhandleMsg, p...)
	nblocks := len(msg) / sm3.BlockSize()
	sm3.update3(msg)
	// Update unhandleMsg
	sm3.unhandleMsg = msg[nblocks*sm3.BlockSize():]
	return toWrite, nil
}

// 把当前 hash 追加到 in 的后面
// 不会改变当前 hash 状态
func (sm3 *SM3) Sum(in []byte) []byte {
	_, _ = sm3.Write(in)
	msg := sm3.pad()
	//FinalizeF
	digest := sm3.update4(msg)

	// save hash to in
	needed := sm3.Size()
	if cap(in)-len(in) < needed {
		newIn := make([]byte, len(in), len(in)+needed)
		copy(newIn, in)
		in = newIn
	}
	out := in[len(in) : len(in)+needed]
	for i := 0; i < 8; i++ {
		binary.BigEndian.PutUint32(out[i*4:], digest[i])
	}
	return out
}

func Sm3Sum(data []byte) []byte {
	var sm3 SM3

	sm3.Reset()
	_, _ = sm3.Write(data)
	return sm3.Sum(nil)
}

//@brief：耗时统计函数
func timeCost(s string) func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf(s+": time cost = %v\n", tc)
	}
}
