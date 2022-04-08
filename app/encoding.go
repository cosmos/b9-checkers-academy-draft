package app

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/tendermint/spm/cosmoscmd"
	appparams "github.com/xavierlepretre/checkers/app/params"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing.
// This function should be used only internally (in the SDK).
// App user should'nt create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() cosmoscmd.EncodingConfig {
	encodingConfig := appparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
