package blog

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/keeper"
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateTitle:
			return handleMsgCreateTitle(ctx, k, msg)
		case types.MsgSetTitle:
			return handleMsgSetTitle(ctx, k, msg)
		case types.MsgDeleteTitle:
			return handleMsgDeleteTitle(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
