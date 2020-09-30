package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteTitle{}

type MsgDeleteTitle struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteTitle(id string, creator sdk.AccAddress) MsgDeleteTitle {
  return MsgDeleteTitle{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteTitle) Route() string {
  return RouterKey
}

func (msg MsgDeleteTitle) Type() string {
  return "DeleteTitle"
}

func (msg MsgDeleteTitle) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteTitle) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteTitle) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}