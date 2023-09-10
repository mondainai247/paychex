package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendPayroll = "send_payroll"

var _ sdk.Msg = &MsgSendPayroll{}

func NewMsgSendPayroll(creator string, newpayroll *Payroll) *MsgSendPayroll {
	return &MsgSendPayroll{
		Creator:    creator,
		Newpayroll: newpayroll,
	}
}

func (msg *MsgSendPayroll) Route() string {
	return RouterKey
}

func (msg *MsgSendPayroll) Type() string {
	return TypeMsgSendPayroll
}

func (msg *MsgSendPayroll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendPayroll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendPayroll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
