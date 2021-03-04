package vm

import (
	"math/big"
	"testing"

	"github.com/abeychain/go-abey/common"
	"github.com/abeychain/go-abey/core/state"
	"github.com/abeychain/go-abey/core/types"
	"github.com/abeychain/go-abey/crypto"
	"github.com/abeychain/go-abey/abeydb"
	"github.com/abeychain/go-abey/log"
	"github.com/abeychain/go-abey/params"
)

func TestDeposit(t *testing.T) {

	priKey, _ := crypto.GenerateKey()
	from := crypto.PubkeyToAddress(priKey.PublicKey)
	pub := crypto.FromECDSAPub(&priKey.PublicKey)
	value := big.NewInt(1000)

	statedb, _ := state.New(common.Hash{}, state.NewDatabase(abeydb.NewMemDatabase()))
	statedb.GetOrNewStateObject(types.StakingAddress)
	evm := NewEVM(Context{}, statedb, params.TestChainConfig, Config{})

	log.Info("Staking deposit", "address", from, "value", value)
	impawn := NewImpawnImpl()
	impawn.Load(evm.StateDB, types.StakingAddress)

	impawn.InsertSAccount2(1000, 0, from, pub, value, big.NewInt(0), true)
	impawn.Save(evm.StateDB, types.StakingAddress)

	impawn1 := NewImpawnImpl()
	impawn1.Load(evm.StateDB, types.StakingAddress)
}
