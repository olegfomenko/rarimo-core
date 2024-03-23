package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "cbank/Deposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "cbank/Withdraw", nil)
	cdc.RegisterConcrete(&MsgTransfer{}, "cbank/Transfer", nil)
	cdc.RegisterConcrete(&MsgFinishTransfer{}, "cbank/FinishTransfer", nil)
	cdc.RegisterConcrete(&Commitment{}, "cbank/Commitment", nil)
	cdc.RegisterConcrete(&RangeProof{}, "cbank/RangeProof", nil)
	cdc.RegisterConcrete(&Signature{}, "cbank/Signature", nil)
	cdc.RegisterConcrete(&Point{}, "cbank/Point", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinishTransfer{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
