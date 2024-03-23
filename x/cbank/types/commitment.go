package types

import (
	"errors"
	"github.com/cloudflare/bn256"
	bulletproofs "github.com/distributed-lab/bulletproofs"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const numBytes = 32

func (p Point) Compress() string {
	return hexutil.Encode(p.CompressBytes())
}

func (p Point) CompressBytes() []byte {
	buf := append(hexutil.MustDecode(p.XCord), hexutil.MustDecode(p.YCord)...)
	if len(buf) != 2*numBytes {
		panic(errors.New("invalid compressed point size"))
	}

	return buf
}

func (p *Point) Decompress(s string) error {
	return p.DecompressBytes(hexutil.MustDecode(s))
}

func (p *Point) DecompressBytes(buf []byte) error {
	if len(buf) != 2*numBytes {
		return errors.New("invalid compressed point size")
	}

	p.XCord = hexutil.Encode(buf[:32])
	p.YCord = hexutil.Encode(buf[32:])
	return nil
}

func (p Point) MustToBN256G1() *bn256.G1 {
	point := new(bn256.G1)
	if _, err := point.Unmarshal(p.CompressBytes()); err != nil {
		panic(err)
	}

	return point
}

func (p *Point) MustFromBN256G1(point *bn256.G1) *Point {
	if err := p.DecompressBytes(point.Marshal()); err != nil {
		panic(err)
	}
	return p
}

type PointSlice []Point

func (p PointSlice) MustToBN256G1() []*bn256.G1 {
	res := make([]*bn256.G1, 0, len(p))

	for _, point := range p {
		res = append(res, point.MustToBN256G1())
	}

	return res
}

func (p *PointSlice) MustFromBN256G1(slice []*bn256.G1) *PointSlice {
	for _, point := range slice {
		(*p) = append((*p), *new(Point).MustFromBN256G1(point))
	}
	return p
}

type ScalarSlice []string

func (s ScalarSlice) MustToBig() []*big.Int {
	res := make([]*big.Int, 0, len(s))

	for _, scalar := range s {
		res = append(res, new(big.Int).SetBytes(hexutil.MustDecode(scalar)))
	}

	return res
}

func (p *RangeProof) Proof() *bulletproofs.ReciprocalProof {
	return &bulletproofs.ReciprocalProof{
		ArithmeticCircuitProof: &bulletproofs.ArithmeticCircuitProof{
			CL: p.Cl.MustToBN256G1(),
			CR: p.Cr.MustToBN256G1(),
			CO: p.Co.MustToBN256G1(),
			CS: p.Cs.MustToBN256G1(),
			WNLA: &bulletproofs.WeightNormLinearArgumentProof{
				R: PointSlice(p.R).MustToBN256G1(),
				X: PointSlice(p.X).MustToBN256G1(),
				L: ScalarSlice(p.L).MustToBig(),
				N: ScalarSlice(p.N).MustToBig(),
			},
		},
		V: p.V.MustToBN256G1(),
	}
}
