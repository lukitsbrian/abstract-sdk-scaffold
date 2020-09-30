package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/keeper"
)

// Handle a message to delete name
func handleMsgDeleteTitle(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteTitle) (*sdk.Result, error) {
	if !k.TitleExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetTitleOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteTitle(ctx, msg.ID)
	return &sdk.Result{}, nil
}
