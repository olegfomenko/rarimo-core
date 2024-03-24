package types

import "github.com/distributed-lab/bulletproofs"

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

func (p Params) ReciprocalPublic() *bulletproofs.ReciprocalPublic {
	return &bulletproofs.ReciprocalPublic{
		G:     p.G.MustToBN256G1(),
		GVec:  PointSlice(p.GVec).MustToBN256G1(),
		HVec:  PointSlice(p.HVec).MustToBN256G1(),
		Nd:    int(p.Nd),
		Np:    int(p.Np),
		GVec_: PointSlice(p.GVecWNLA).MustToBN256G1(),
		HVec_: PointSlice(p.HVecWNLA).MustToBN256G1(),
	}
}
