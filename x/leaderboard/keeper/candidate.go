package keeper

import (
	"github.com/b9lab/checkers/x/leaderboard/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetCandidate(ctx sdk.Context, candidate types.Candidate) {
	candidateStore := prefix.NewStore(ctx.TransientStore(k.tKey), []byte(types.CandidateKeyPrefix))
	candidateBytes := k.cdc.MustMarshal(&candidate)
	candidateStore.Set(candidate.Address, candidateBytes)
}

func (k Keeper) GetAllCandidates(ctx sdk.Context) (candidates []types.Candidate) {
	candidateStore := prefix.NewStore(ctx.TransientStore(k.tKey), []byte(types.CandidateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(candidateStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var candidate types.Candidate
		k.cdc.MustUnmarshal(iterator.Value(), &candidate)
		candidates = append(candidates, candidate)
	}

	return
}
