package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DistributeFeePool distributes funds from the the community pool to a receiver address
func (k Keeper) DistributeFeePool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error {
	feePool := k.GetFeePool(ctx)

	poolTruncted, _ := feePool.CommunityPool.TruncateDecimal()
	if !poolTruncted.IsAllGTE(amount) {
		return errors.New("community pool does not have sufficient coins to distribute")
	}

	feePool.CommunityPool.Sub(sdk.NewDecCoins(amount))
	k.bankKeeper.AddCoins(ctx, receiveAddr, amount)

	k.SetFeePool(ctx, feePool)
	return nil
}