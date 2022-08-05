package precompile

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gonum.org/v1/gonum/stat"
)

type ContractMomentConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractMomentConfig) Address() common.Address {
	return ContractMomentAddress
}

func (c ContractMomentConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractMomentConfig) IsDisabled() bool {
	return false
}

func (c ContractMomentConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractMomentConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractMomentConfig) Contract() StatefulPrecompiledContract {
	return ContractMomentPrecompile
}

var (
	ContractMomentPrecompile StatefulPrecompiledContract = createMomentPrecompile(ContractMomentAddress)
	momentSignature                                      = crypto.Keccak256([]byte("getMoment(uint256,uint256[],uint256[])"))[:4]
)

func mustMomentType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeMomentArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "v1",
			Type: mustMomentType("uint256"),
		},
		{
			Name: "v2",
			Type: mustMomentType("uint256[]"),
		},
		{
			Name: "v3",
			Type: mustMomentType("uint256[]"),
		},
	}
}

func MakeMomentRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func bigIntToFloat64(v []*big.Int) []float64 {
	var output []float64
	for i := range v {
		f := new(big.Float).SetInt(v[i])
		s, _ := f.Float64()
		output = append(output, s)
	}
	return output
}

func calcMoment(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	vals, err := MakeMomentArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	if len(vals) != 3 {
		return nil, suppliedGas, errors.New("Invalid number of args")
	}

	v1, ok := vals[0].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	f := new(big.Float).SetInt(v1)
	moment, _ := f.Float64()

	v2, ok := vals[1].([]*big.Int)
	v3, ok := vals[2].([]*big.Int)

	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}

	samples := bigIntToFloat64(v2)
	weights := bigIntToFloat64(v3)

	_result := stat.Moment(moment, samples, weights)

	bigval := new(big.Float).SetFloat64(_result)
	rr, _ := bigval.Int64()
	result := new(big.Int).SetInt64(rr)

	ret, err = MakeMomentRetArgs().PackValues([]interface{}{result})

	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}

func createMomentPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(momentSignature, calcMoment)})
}
