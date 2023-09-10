package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"paychex/x/paychex/types"
)

// GetEmployeeCount get the total number of employee
func (k Keeper) GetEmployeeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EmployeeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEmployeeCount set the total number of employee
func (k Keeper) SetEmployeeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EmployeeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEmployee appends a employee in the store with a new id and update the count
func (k Keeper) AppendEmployee(
	ctx sdk.Context,
	employee types.Employee,
) uint64 {
	// Create the employee
	count := k.GetEmployeeCount(ctx)

	// Set the ID of the appended value
	employee.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployeeKey))
	appendedValue := k.cdc.MustMarshal(&employee)
	store.Set(GetEmployeeIDBytes(employee.Id), appendedValue)

	// Update employee count
	k.SetEmployeeCount(ctx, count+1)

	return count
}

// SetEmployee set a specific employee in the store
func (k Keeper) SetEmployee(ctx sdk.Context, employee types.Employee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployeeKey))
	b := k.cdc.MustMarshal(&employee)
	store.Set(GetEmployeeIDBytes(employee.Id), b)
}

// GetEmployee returns a employee from its id
func (k Keeper) GetEmployee(ctx sdk.Context, id uint64) (val types.Employee, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployeeKey))
	b := store.Get(GetEmployeeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEmployee removes a employee from the store
func (k Keeper) RemoveEmployee(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployeeKey))
	store.Delete(GetEmployeeIDBytes(id))
}

// GetAllEmployee returns all employee
func (k Keeper) GetAllEmployee(ctx sdk.Context) (list []types.Employee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployeeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Employee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEmployeeIDBytes returns the byte representation of the ID
func GetEmployeeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEmployeeIDFromBytes returns ID in uint64 format from a byte array
func GetEmployeeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
