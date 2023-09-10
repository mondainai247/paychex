package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateEmployee{}, "paychex/CreateEmployee", nil)
	cdc.RegisterConcrete(&MsgUpdateEmployee{}, "paychex/UpdateEmployee", nil)
	cdc.RegisterConcrete(&MsgDeleteEmployee{}, "paychex/DeleteEmployee", nil)
	cdc.RegisterConcrete(&MsgCreateQueryResponse{}, "paychex/CreateQueryResponse", nil)
	cdc.RegisterConcrete(&MsgUpdateQueryResponse{}, "paychex/UpdateQueryResponse", nil)
	cdc.RegisterConcrete(&MsgDeleteQueryResponse{}, "paychex/DeleteQueryResponse", nil)
	cdc.RegisterConcrete(&MsgSendPayroll{}, "paychex/SendPayroll", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEmployee{},
		&MsgUpdateEmployee{},
		&MsgDeleteEmployee{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateQueryResponse{},
		&MsgUpdateQueryResponse{},
		&MsgDeleteQueryResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendPayroll{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
