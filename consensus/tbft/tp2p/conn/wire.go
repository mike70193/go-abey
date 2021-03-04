package conn

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/abeychain/go-abey/consensus/tbft/crypto/cryptoamino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
