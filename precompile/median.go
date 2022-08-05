package precompile

import (
	"errors"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	ContractMedianPrecompile StatefulPrecompiledContract = createMedianPrecompile(ContractMedianAddress)
	medianSignature                                      = crypto.Keccak256([]byte("getMedian(uint256[])"))[:4]
)

func mustType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "vals",
			Type: mustType("uint256[]"),
		},
	}
}

func MakeRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func getMedian(evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	valsI, err := MakeArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	vals, ok := valsI[0].([]*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid type")
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].Cmp(vals[j]) == -1
	})
	med := vals[len(vals)/2]
	ret, err = MakeRetArgs().PackValues([]interface{}{med})
	if err != nil {
		return nil, suppliedGas, err
	}
	return ret, suppliedGas, nil
}

func createMedianPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(medianSignature, getMedian)})
	return contract
}
