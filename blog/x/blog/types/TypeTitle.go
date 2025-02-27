package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Title struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Body string `json:"body" yaml:"body"`
}