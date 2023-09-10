package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEmployee = "create_employee"
	TypeMsgUpdateEmployee = "update_employee"
	TypeMsgDeleteEmployee = "delete_employee"
)

var _ sdk.Msg = &MsgCreateEmployee{}

func NewMsgCreateEmployee(creator string, name string, role *Role) *MsgCreateEmployee {
	return &MsgCreateEmployee{
		Creator: creator,
		Name:    name,
		Role:    role,
	}
}

func (msg *MsgCreateEmployee) Route() string {
	return RouterKey
}

func (msg *MsgCreateEmployee) Type() string {
	return TypeMsgCreateEmployee
}

func (msg *MsgCreateEmployee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEmployee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEmployee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEmployee{}

func NewMsgUpdateEmployee(creator string, id uint64, name string, role *Role) *MsgUpdateEmployee {
	return &MsgUpdateEmployee{
		Id:      id,
		Creator: creator,
		Name:    name,
		Role:    role,
	}
}

func (msg *MsgUpdateEmployee) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEmployee) Type() string {
	return TypeMsgUpdateEmployee
}

func (msg *MsgUpdateEmployee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEmployee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEmployee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEmployee{}

func NewMsgDeleteEmployee(creator string, id uint64) *MsgDeleteEmployee {
	return &MsgDeleteEmployee{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteEmployee) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEmployee) Type() string {
	return TypeMsgDeleteEmployee
}

func (msg *MsgDeleteEmployee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEmployee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEmployee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
