package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateTitle{}

type MsgCreateTitle struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Body string `json:"body" yaml:"body"`
}

func NewMsgCreateTitle(creator sdk.AccAddress, body string) MsgCreateTitle {
  return MsgCreateTitle{
    ID: uuid.New().String(),
		Creator: creator,
    Body: body,
	}
}

func (msg MsgCreateTitle) Route() string {
  return RouterKey
}

func (msg MsgCreateTitle) Type() string {
  return "CreateTitle"
}

func (msg MsgCreateTitle) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateTitle) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateTitle) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}