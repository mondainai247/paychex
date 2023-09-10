package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateQueryResponse = "create_query_response"
	TypeMsgUpdateQueryResponse = "update_query_response"
	TypeMsgDeleteQueryResponse = "delete_query_response"
)

var _ sdk.Msg = &MsgCreateQueryResponse{}

func NewMsgCreateQueryResponse(creator string, staff *Staff) *MsgCreateQueryResponse {
	return &MsgCreateQueryResponse{
		Creator: creator,
		Staff:   staff,
	}
}

func (msg *MsgCreateQueryResponse) Route() string {
	return RouterKey
}

func (msg *MsgCreateQueryResponse) Type() string {
	return TypeMsgCreateQueryResponse
}

func (msg *MsgCreateQueryResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateQueryResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateQueryResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateQueryResponse{}

func NewMsgUpdateQueryResponse(creator string, id uint64, staff *Staff) *MsgUpdateQueryResponse {
	return &MsgUpdateQueryResponse{
		Id:      id,
		Creator: creator,
		Staff:   staff,
	}
}

func (msg *MsgUpdateQueryResponse) Route() string {
	return RouterKey
}

func (msg *MsgUpdateQueryResponse) Type() string {
	return TypeMsgUpdateQueryResponse
}

func (msg *MsgUpdateQueryResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateQueryResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateQueryResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteQueryResponse{}

func NewMsgDeleteQueryResponse(creator string, id uint64) *MsgDeleteQueryResponse {
	return &MsgDeleteQueryResponse{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteQueryResponse) Route() string {
	return RouterKey
}

func (msg *MsgDeleteQueryResponse) Type() string {
	return TypeMsgDeleteQueryResponse
}

func (msg *MsgDeleteQueryResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteQueryResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteQueryResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
