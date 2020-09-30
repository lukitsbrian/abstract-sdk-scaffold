package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetTitle{}

type MsgSetTitle struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Body string `json:"body" yaml:"body"`
}

func NewMsgSetTitle(creator sdk.AccAddress, id string, body string) MsgSetTitle {
  return MsgSetTitle{
    ID: id,
		Creator: creator,
    Body: body,
	}
}

func (msg MsgSetTitle) Route() string {
  return RouterKey
}

func (msg MsgSetTitle) Type() string {
  return "SetTitle"
}

func (msg MsgSetTitle) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetTitle) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetTitle) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}