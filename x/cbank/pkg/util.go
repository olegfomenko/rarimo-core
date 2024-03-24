package pkg

import (
	"crypto/rand"
	"github.com/cloudflare/bn256"
	"math/big"
)

func PublicKey(prv *big.Int, g *bn256.G1) *bn256.G1 {
	return new(bn256.G1).ScalarMult(g, prv)
}

func PrivateKey(g *bn256.G1) (*big.Int, *bn256.G1, error) {
	r, err := rand.Int(rand.Reader, bn256.Order)
	if err != nil {
		return nil, nil, ErrFailedRandom
	}

	return r, new(bn256.G1).ScalarMult(g, r), nil
}
