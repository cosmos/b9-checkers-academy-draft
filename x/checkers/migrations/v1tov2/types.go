package v1tov2

import "github.com/b9lab/checkers/x/checkers/types"

type GenesisStateV1 struct {
	Params         types.Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SystemInfo     types.SystemInfo   `protobuf:"bytes,2,opt,name=systemInfo,proto3" json:"systemInfo"`
	StoredGameList []types.StoredGame `protobuf:"bytes,3,rep,name=storedGameList,proto3" json:"storedGameList"`
}
