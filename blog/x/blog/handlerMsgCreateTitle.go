package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/keeper"
)

func handleMsgCreateTitle(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateTitle) (*sdk.Result, error) {
	var title = types.Title{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Body: msg.Body,
	}
	k.CreateTitle(ctx, title)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
