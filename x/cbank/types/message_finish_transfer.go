package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgFinishTransfer = "finish-transfer"
)

var _ sdk.Msg = &MsgFinishTransfer{}

func (msg *MsgFinishTransfer) Route() string {
	return RouterKey
}

func (msg *MsgFinishTransfer) Type() string {
	return TypeMsgFinishTransfer
}

func (msg *MsgFinishTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinishTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinishTransfer) ValidateBasic() error {
	return nil
}
