package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"paychex/x/paychex/types"
)

// GetQueryResponseCount get the total number of queryResponse
func (k Keeper) GetQueryResponseCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.QueryResponseCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetQueryResponseCount set the total number of queryResponse
func (k Keeper) SetQueryResponseCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.QueryResponseCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendQueryResponse appends a queryResponse in the store with a new id and update the count
func (k Keeper) AppendQueryResponse(
	ctx sdk.Context,
	queryResponse types.QueryResponse,
) uint64 {
	// Create the queryResponse
	count := k.GetQueryResponseCount(ctx)

	// Set the ID of the appended value
	queryResponse.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QueryResponseKey))
	appendedValue := k.cdc.MustMarshal(&queryResponse)
	store.Set(GetQueryResponseIDBytes(queryResponse.Id), appendedValue)

	// Update queryResponse count
	k.SetQueryResponseCount(ctx, count+1)

	return count
}

// SetQueryResponse set a specific queryResponse in the store
func (k Keeper) SetQueryResponse(ctx sdk.Context, queryResponse types.QueryResponse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QueryResponseKey))
	b := k.cdc.MustMarshal(&queryResponse)
	store.Set(GetQueryResponseIDBytes(queryResponse.Id), b)
}

// GetQueryResponse returns a queryResponse from its id
func (k Keeper) GetQueryResponse(ctx sdk.Context, id uint64) (val types.QueryResponse, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QueryResponseKey))
	b := store.Get(GetQueryResponseIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveQueryResponse removes a queryResponse from the store
func (k Keeper) RemoveQueryResponse(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QueryResponseKey))
	store.Delete(GetQueryResponseIDBytes(id))
}

// GetAllQueryResponse returns all queryResponse
func (k Keeper) GetAllQueryResponse(ctx sdk.Context) (list []types.QueryResponse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QueryResponseKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.QueryResponse
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetQueryResponseIDBytes returns the byte representation of the ID
func GetQueryResponseIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetQueryResponseIDFromBytes returns ID in uint64 format from a byte array
func GetQueryResponseIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
