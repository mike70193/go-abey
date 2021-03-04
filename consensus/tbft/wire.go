package tbft

import (
	"github.com/tendermint/go-amino"
	"github.com/abeychain/go-abey/consensus/tbft/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterConsensusMessages(cdc)
	// RegisterWALMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
