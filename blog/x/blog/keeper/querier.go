package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for blog clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListTitle:
			return listTitle(ctx, k)
		case types.QueryGetTitle:
			return getTitle(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown blog query endpoint")
		}
	}
}
