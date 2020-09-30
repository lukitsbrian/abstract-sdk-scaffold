package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateTitle creates a title
func (k Keeper) CreateTitle(ctx sdk.Context, title types.Title) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.TitlePrefix + title.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(title)
	store.Set(key, value)
}

// GetTitle returns the title information
func (k Keeper) GetTitle(ctx sdk.Context, key string) (types.Title, error) {
	store := ctx.KVStore(k.storeKey)
	var title types.Title
	byteKey := []byte(types.TitlePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &title)
	if err != nil {
		return title, err
	}
	return title, nil
}

// SetTitle sets a title
func (k Keeper) SetTitle(ctx sdk.Context, title types.Title) {
	titleKey := title.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(title)
	key := []byte(types.TitlePrefix + titleKey)
	store.Set(key, bz)
}

// DeleteTitle deletes a title
func (k Keeper) DeleteTitle(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.TitlePrefix + key))
}

//
// Functions used by querier
//

func listTitle(ctx sdk.Context, k Keeper) ([]byte, error) {
	var titleList []types.Title
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TitlePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var title types.Title
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &title)
		titleList = append(titleList, title)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, titleList)
	return res, nil
}

func getTitle(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	title, err := k.GetTitle(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, title)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetTitleOwner(ctx sdk.Context, key string) sdk.AccAddress {
	title, err := k.GetTitle(ctx, key)
	if err != nil {
		return nil
	}
	return title.Creator
}


// Check if the key exists in the store
func (k Keeper) TitleExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.TitlePrefix + key))
}
