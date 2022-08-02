package precompile

import (
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	medianSignature                                      = CalculateFunctionSelector("getMedian(uint256,uint256,uint256)")
)

func mustType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "v1",
			Type: mustType("uint256"),
		},
		{
			Name: "v2",
			Type: mustType("uint256"),
		},
		{
			Name: "v3",
			Type: mustType("uint256"),
		},
	}
}

func MakeRetArgs() abi.Arguments {
	//cocos,taste of thai express,sangcan indian
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func createMedianPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	f := func(
		evm PrecompileAccessibleState,
		callerAddr common.Address,
		addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
		//var b [32]byte
		//b[31] = 0xaa
		//return b[:], suppliedGas, nil
		//log.Info("Info:", hexutil.Encode(input))
		fmt.Println("input", hexutil.Encode(input))
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)
		var errb [32]byte
		errb[31] = 0xaa
		vals, err := MakeArgs().UnpackValues(inputCopy)
		if err != nil {
			return errb[:], suppliedGas, err
		}
		if len(vals) != 3 {
			return errb[:], suppliedGas, errors.New("invalid vals")
		}
		v1, ok := vals[0].(*big.Int)
		if !ok {
			return errb[:], suppliedGas, errors.New("invalid val")
		}
		v2, ok := vals[1].(*big.Int)
		if !ok {
			return nil, suppliedGas, errors.New("invalid val")
		}
		v3, ok := vals[2].(*big.Int)
		if !ok {
			return nil, suppliedGas, errors.New("invalid val")
		}
		valsI := []*big.Int{v1, v2, v3}
		sort.Slice(valsI, func(i, j int) bool {
			return valsI[i].Cmp(valsI[j]) == -1
		})
		med := valsI[len(valsI)/2]
		ret, err = MakeRetArgs().PackValues([]interface{}{med})
		if err != nil {
			return errb[:], suppliedGas, err
		}
		return ret, suppliedGas, nil
	}
	funcGetXChainECRecover := newStatefulPrecompileFunction(medianSignature, f)
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{funcGetXChainECRecover})
	return contract
}
