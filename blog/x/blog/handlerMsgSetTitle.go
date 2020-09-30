package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/keeper"
)

func handleMsgSetTitle(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetTitle) (*sdk.Result, error) {
	var title = types.Title{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Body: msg.Body,
	}
	if !msg.Creator.Equals(k.GetTitleOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetTitle(ctx, title)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
