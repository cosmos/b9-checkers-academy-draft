package keeper

type Hooks struct {
	k Keeper
}

func (k Keeper) Hooks() Hooks { return Hooks{k} }
