package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgTransfer = "transfer"
)

var _ sdk.Msg = &MsgTransfer{}

func (msg *MsgTransfer) Route() string {
	return RouterKey
}

func (msg *MsgTransfer) Type() string {
	return TypeMsgTransfer
}

func (msg *MsgTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// TODO validation + denom check
func (msg *MsgTransfer) ValidateBasic() error {
	return nil
}
