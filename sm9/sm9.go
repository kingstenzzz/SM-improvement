// Copyright 2020 cetc-30. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm9

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"math"
	"math/big"

	"github.com/kingstenzzz/SM-improvement/sm3"
	"github.com/kingstenzzz/SM-improvement/sm9/sm9curve"
	"github.com/pkg/errors"
)

type hashMode int

const (
	H1 hashMode = iota
	H2
)

//MasterKey contains a master secret key and a master public key.
type MasterKey struct {
	Msk *big.Int
	MasterPubKey
}

type MasterPubKey struct {
	Mpk *sm9curve.G2
}

//UserKey contains a secret key.
type UserKey struct {
	Sk *sm9curve.G1
}

//Sm9Sig contains a big number and an element in G1.
type Sm9Sig struct {
	H *big.Int
	S *sm9curve.G1
}

//hash implements H1(Z,n) or H2(Z,n) in sm9 algorithm.
//z为比特串，n为整数，输出结果范围是[1,n-1]
//在签名的生成和验证过程之前，应用密码杂凑函数对待签消息M和待验证消息M′进行杂凑计算。
func hash(z []byte, n *big.Int, h hashMode) *big.Int {
	//counter
	ct := 1 //初始化32bit构成的计数器
	//Ceil是向上取值
	hlen := 8 * int(math.Ceil(float64(5*n.BitLen()/32)))

	var ha []byte
	//哈希值为256位
	for i := 0; i < int(math.Ceil(float64(hlen/256))); i++ {
		msg := append([]byte{byte(h)}, z...)
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(ct)) //将ct转为byte
		msg = append(msg, buf...)
		hai := sm3.Sm3Sum(msg)
		ct++
		//整数&&
		if float64(hlen)/256 == float64(int64(hlen/256)) &&
			i == int(math.Ceil(float64(hlen/256)))-1 {
			ha = append(ha, hai[:(hlen-256*int(math.Floor(float64(hlen/256))))/32]...)
		} else {
			ha = append(ha, hai[:]...)
		}
	}
	//步骤f
	bn := new(big.Int).SetBytes(ha)
	one := big.NewInt(1)
	nMinus1 := new(big.Int).Sub(n, one)
	bn.Mod(bn, nMinus1)
	bn.Add(bn, one)

	return bn
}

//随机数发生器
//generate rand numbers in [1,n-1].
func randFieldElement(rand io.Reader, n *big.Int) (k *big.Int, err error) {
	one := big.NewInt(1)
	b := make([]byte, 256/8+8) //32+8
	_, err = io.ReadFull(rand, b)
	if err != nil {
		return
	}
	k = new(big.Int).SetBytes(b)
	nMinus1 := new(big.Int).Sub(n, one)
	k.Mod(k, nMinus1)
	return
}

//generate master key for KGC(Key Generate Center).
//主公私钥对
//签名主私钥一般由ＫＧＣ通过随机数发生器产生，签名主公钥由签名主私钥结合系统参数产生
func MasterKeyGen(rand io.Reader) (mk *MasterKey, err error) {
	s, err := randFieldElement(rand, sm9curve.Order)
	if err != nil {
		return nil, errors.Errorf("gen rand num err:%s", err)
	}

	mk = new(MasterKey)
	mk.Msk = new(big.Int).Set(s)

	mk.Mpk = new(sm9curve.G2).ScalarBaseMult(s) //[ks]P2

	return
}

//产生用户签名密钥
//generate user's secret key.
func UserKeyGen(mk *MasterKey, id []byte, hid byte) (uk *UserKey, err error) {
	id = append(id, hid) //hid为用一个字节表示的签名私钥生成函数识别符
	n := sm9curve.Order
	t1 := hash(id, n, H1)
	t1.Add(t1, mk.Msk)

	//if t1 = 0, we need to regenerate the master key.
	if t1.BitLen() == 0 || t1.Cmp(n) == 0 {
		return nil, errors.New("need to regen mk!")
	}

	t1.ModInverse(t1, n)

	//t2 = s*t1^-1
	t2 := new(big.Int).Mul(mk.Msk, t1)

	uk = new(UserKey)
	uk.Sk = new(sm9curve.G1).ScalarBaseMult(t2)
	return
}

//sm9 sign algorithm:数字签名生成算法流程
//A1:compute g = e(P1,Ppub);
//A2:choose random num r in [1,n-1];
//A3:compute w = g^r;
//A4:compute h = H2(M||w,n);
//A5:compute l = (r-h) mod n, if l = 0 goto A2;
//A6:compute S = l·sk.
func Sign(uk *UserKey, mpk *MasterPubKey, msg []byte) (sig *Sm9Sig, err error) {
	sig = new(Sm9Sig)
	n := sm9curve.Order

	g := sm9curve.Pair(sm9curve.Gen1, mpk.Mpk)
	Gtmp = g
regen:
	r, err := randFieldElement(rand.Reader, n)
	if err != nil {
		return nil, errors.Errorf("gen rand num failed:%s", err)
	}

	w := new(sm9curve.GT).ScalarMult(g, r) //g^r

	wBytes := w.Marshal()

	msg = append(msg, wBytes...)

	h := hash(msg, n, H2)

	sig.H = new(big.Int).Set(h)

	l := new(big.Int).Sub(r, h)
	l.Mod(l, n)

	if l.BitLen() == 0 {
		goto regen
	}

	sig.S = new(sm9curve.G1).ScalarMult(uk.Sk, l)

	return
}

//func NewVerify(sig *Sm9Sig, msg []byte, id []byte, hid byte, mpk *MasterPubKey) bool {}
//sm9 verify algorithm(given sig (h',S'), message M' and user's id):
//B1:compute g = e(P1,Ppub);
//B2:compute t = g^h';
//B3:compute h1 = H1(id||hid,n);
//B4:compute P = h1·P2+Ppub;
//B5:compute u = e(S',P);
//B6:compute w' = u·t;
//B7:compute h2 = H2(M'||w',n), check if h2 = h'.
func Verify(sig *Sm9Sig, msg []byte, id []byte, hid byte, mpk *MasterPubKey) bool {
	n := sm9curve.Order
	g := sm9curve.Pair(sm9curve.Gen1, mpk.Mpk)

	t := new(sm9curve.GT).ScalarMult(g, sig.H)

	id = append(id, hid)

	h1 := hash(id, n, H1)

	P := new(sm9curve.G2).ScalarBaseMult(h1)

	P.Add(P, mpk.Mpk)

	u := sm9curve.Pair(sig.S, P)

	w := new(sm9curve.GT).Add(u, t)

	wBytes := w.Marshal()

	msg = append(msg, wBytes...)

	h2 := hash(msg, n, H2)

	if h2.Cmp(sig.H) != 0 {
		return false
	}

	return true
}

/*func NewVerify(sig *Sm9Sig, msg []byte, id []byte, hid byte, mpk *MasterPubKey) bool {

	n := sm9curve.Order
	var g *sm9curve.GT
	if flag == 0 {
		Gtmp = sm9curve.Pair(sm9curve.Gen1, mpk.Mpk)
		g = Gtmp
		flag = 1
	} else {
		g = Gtmp
	}

	t := new(sm9curve.GT).NewScalarMult(g, g2, g4, g8, sig.H) //t=g^h'

	id = append(id, hid)

	h1 := hash(id, n, H1)

	P := new(sm9curve.G2).ScalarBaseMult(h1)

	P.Add(P, mpk.Mpk)

	u := sm9curve.Pair(sig.S, P)

	w := new(sm9curve.GT).Add(u, t)

	wBytes := w.Marshal()

	msg = append(msg, wBytes...)

	h2 := hash(msg, n, H2)

	if h2.Cmp(sig.H) != 0 {
		return false
	}

	return true
}*/

/*func NewSign(uk *UserKey, mpk *MasterPubKey, msg []byte) (sig *Sm9Sig, err error) {
	sig = new(Sm9Sig)
	n := sm9curve.Order
	var g *sm9curve.GT
	if flag == 0 {
		Gtmp = sm9curve.Pair(sm9curve.Gen1, mpk.Mpk)
		g = Gtmp
		flag = 1
	} else {
		g = Gtmp
	}
regen:
	r, err := randFieldElement(rand.Reader, n)
	if err != nil {
		errors.Errorf("gen rand num failed:%s", err)
	}
	g2 = new(sm9curve.GT).ScalarMult(g, big.NewInt(64))
	g4 = new(sm9curve.GT).ScalarMult(g, big.NewInt(128))
	g8 = new(sm9curve.GT).ScalarMult(g, big.NewInt(192))
	w := new(sm9curve.GT).NewScalarMult(g, g2, g4, g8, r)
	w2 := new(sm9curve.GT).ScalarMult(g, r)
	wBytes := w.Marshal()
	w2Bytes := w2.Marshal()
	fmt.Println("w1:", wBytes)
	fmt.Println("w2:", w2Bytes)
	msg = append(msg, wBytes...)

	h := hash(msg, n, H2)

	sig.H = new(big.Int).Set(h)

	l := new(big.Int).Sub(r, h)
	l.Mod(l, n)

	if l.BitLen() == 0 {
		goto regen
	}

	sig.S = new(sm9curve.G1).ScalarMult(uk.Sk, l)

	return

}*/
