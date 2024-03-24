package types

import (
	"errors"
	"github.com/cloudflare/bn256"
	bulletproofs "github.com/distributed-lab/bulletproofs"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/rarimo-core/x/cbank/pkg"
	"math/big"
)

const numBytes = 32

func NewCommitment(com *bn256.G1, addr *bn256.G1, denom string) Commitment {
	return Commitment{
		Commitment: *new(Point).MustFromBN256G1(com),
		Address:    *new(Point).MustFromBN256G1(addr),
		Denom:      denom,
	}
}

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

func (p *Point) Decompress(s string) (*Point, error) {
	return p, p.DecompressBytes(hexutil.MustDecode(s))
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

func (s *ScalarSlice) MustFromBig(arr []*big.Int) *ScalarSlice {
	for _, b := range arr {
		*s = append(*s, hexutil.Encode(b.Bytes()))
	}

	return s
}

func (c *Commitment) Index() string {
	return hexutil.Encode(new(bn256.G1).Add(c.Commitment.MustToBN256G1(), c.Address.MustToBN256G1()).Marshal())
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

func (p *RangeProof) FromProof(proof *bulletproofs.ReciprocalProof) *RangeProof {
	return &RangeProof{
		Cl: *new(Point).MustFromBN256G1(proof.CL),
		Cr: *new(Point).MustFromBN256G1(proof.CR),
		Co: *new(Point).MustFromBN256G1(proof.CO),
		Cs: *new(Point).MustFromBN256G1(proof.CS),
		X:  *new(PointSlice).MustFromBN256G1(proof.WNLA.X),
		R:  *new(PointSlice).MustFromBN256G1(proof.WNLA.R),
		L:  *new(ScalarSlice).MustFromBig(proof.WNLA.L),
		N:  *new(ScalarSlice).MustFromBig(proof.WNLA.N),
		V:  *new(Point).MustFromBN256G1(proof.V),
	}

}

func (s *Signature) Signature() *pkg.SchnorrSignature {
	return &pkg.SchnorrSignature{
		R: s.R.MustToBN256G1(),
		S: new(big.Int).SetBytes(hexutil.MustDecode(s.S)),
	}
}

func (s *Signature) FromSignature(sig *pkg.SchnorrSignature) *Signature {
	return &Signature{
		R: *new(Point).MustFromBN256G1(sig.R),
		S: hexutil.Encode(sig.S.Bytes()),
	}
}
