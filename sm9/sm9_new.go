package sm9

import (
	"crypto/rand"

	"math/big"

	"github.com/kingstenzzz/SM-improvemnt/sm9/sm9curve"
	"github.com/pkg/errors"
)

var Gtmp *sm9curve.GT

var Mpk *sm9curve.G2

func NewSign(uk *UserKey, mpk *MasterPubKey, msg []byte) (sig *Sm9Sig, err error) {

	sig = new(Sm9Sig)
	n := sm9curve.Order
	var g *sm9curve.GT
	if Mpk != mpk.Mpk {
		Gtmp = sm9curve.Pair(sm9curve.Gen1, mpk.Mpk)
		g = Gtmp
		Mpk = mpk.Mpk
	} else {
		g = Gtmp
	}

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

func NewVerify(sig *Sm9Sig, msg []byte, id []byte, hid byte, mpk *MasterPubKey) bool {
	var g *sm9curve.GT
	n := sm9curve.Order
	g = Gtmp

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
