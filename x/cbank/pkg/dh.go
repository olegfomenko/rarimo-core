package pkg

import (
	"github.com/cloudflare/bn256"
	"math/big"
)

func HashDH(publicKey *bn256.G1, privateKey *big.Int) *big.Int {
	return new(big.Int).Mod(
		new(big.Int).SetBytes(new(bn256.G1).ScalarMult(publicKey, privateKey).Marshal()),
		bn256.Order,
	)
}
