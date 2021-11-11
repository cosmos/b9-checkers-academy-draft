package v1tov2

import (
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

type GenesisStateV1 struct {
	StoredGameList []*types.StoredGame `protobuf:"bytes,2,rep,name=storedGameList,proto3" json:"storedGameList,omitempty"`
	NextGame       *types.NextGame     `protobuf:"bytes,1,opt,name=nextGame,proto3" json:"nextGame,omitempty"`
}
