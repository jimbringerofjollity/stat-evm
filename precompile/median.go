package precompile

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ContractMedianConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractMedianConfig) Address() common.Address {
	return ContractMedianAddress
}

func (c ContractMedianConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractMedianConfig) IsDisabled() bool {
	return false
}

func (c ContractMedianConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractMedianConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractMedianConfig) Contract() StatefulPrecompiledContract {
	return ContractMedianPrecompile
}

var (
	_ StatefulPrecompileConfig = &ContractXChainECRecoverConfig{}
	// Singleton StatefulPrecompiledContract for XChain ECRecover.
	ContractMedianPrecompile StatefulPrecompiledContract = createMedianPrecompile(ContractMedianAddress)
	medianSignature                                      = CalculateFunctionSelector("getMedian()")
)

func createMedianPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	f := func(
		evm PrecompileAccessibleState,
		callerAddr common.Address,
		addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
		var b [32]byte
		b[28] = 0x0b
		b[29] = 0x0e
		b[30] = 0x0e
		b[31] = 0x0f
		return b[:], suppliedGas, nil
	}
	funcGetXChainECRecover := newStatefulPrecompileFunction(medianSignature, f)
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{funcGetXChainECRecover})
	return contract
}
