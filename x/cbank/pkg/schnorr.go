package pkg

import (
	"bytes"
	"crypto/rand"
	"errors"
	"github.com/cloudflare/bn256"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type SchnorrSignature struct {
	R *bn256.G1
	S *big.Int
}

type HashF func(bytes ...[]byte) []byte

var Hash HashF = crypto.Keccak256

var ErrFailedRandom = errors.New("failed to generate secure random")

func R(G *bn256.G1) (*big.Int, *bn256.G1, error) {
	r, err := rand.Int(rand.Reader, bn256.Order)
	if err != nil {
		return nil, nil, ErrFailedRandom
	}

	return r, new(bn256.G1).ScalarMult(G, r), nil
}

func Msg(bytes ...[]byte) *big.Int {
	return new(big.Int).Mod(new(big.Int).SetBytes(Hash(bytes...)), bn256.Order)
}

func MultiSigSchnorr(prv *big.Int, r *big.Int, PubKeyCommon *bn256.G1, RCommon *bn256.G1, m *big.Int) (*SchnorrSignature, error) {
	hash := Msg(m.Bytes(), PubKeyCommon.Marshal(), RCommon.Marshal())
	s := new(big.Int).Add(r, new(big.Int).Mul(hash, prv))

	return &SchnorrSignature{
		R: RCommon,
		S: s,
	}, nil
}

func SignSchnorr(prv *big.Int, PublicKey *bn256.G1, G *bn256.G1, m *big.Int) (*SchnorrSignature, error) {
	r, R, err := R(G)
	if err != nil {
		return nil, err
	}

	hash := Msg(m.Bytes(), PublicKey.Marshal(), R.Marshal())

	s := new(big.Int).Add(r, new(big.Int).Mul(hash, prv))

	return &SchnorrSignature{
		R: R,
		S: s,
	}, nil
}

func VerifySchnorr(sig *SchnorrSignature, PublicKey *bn256.G1, G *bn256.G1, m *big.Int) bool {
	// s = r + hash*prv
	// R = rG

	// p2 = (r + hash*prv)*G
	// p1 = rG + hash*prv*G = (r + hash*prv)*G

	hash := Msg(m.Bytes(), PublicKey.Marshal(), sig.R.Marshal())

	p1 := new(bn256.G1).ScalarMult(PublicKey, hash)
	p1 = new(bn256.G1).Add(p1, sig.R)

	p2 := new(bn256.G1).ScalarMult(G, sig.S)

	return bytes.Equal(p1.Marshal(), p2.Marshal())
}
