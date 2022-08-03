package precompile

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type ContractFitConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractFitConfig) Address() common.Address {
	return ContractFitAddress
}

func (c ContractFitConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractFitConfig) IsDisabled() bool {
	return false
}

func (c ContractFitConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractFitConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractFitConfig) Contract() StatefulPrecompiledContract {
	return ContractFitPrecompile
}

var (
	_ StatefulPrecompileConfig = &ContractXChainECRecoverConfig{}
	// Singleton StatefulPrecompiledContract for XChain ECRecover.
	ContractFitPrecompile StatefulPrecompiledContract = createFitPrecompile(ContractFitAddress)
	fitSignature                                      = CalculateFunctionSelector("getFit(uint256,uint256,uint256,uint256,uint256,uint256)")
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
		{
			Name: "v4",
			Type: mustType("uint256"),
		},
		{
			Name: "v5",
			Type: mustType("uint256"),
		},
		{
			Name: "v6",
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

func createFitPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	f := func(
		evm PrecompileAccessibleState,
		callerAddr common.Address,
		addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (retb []byte, remainingGas uint64, err error) {
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
		if len(vals) != 6 {
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
		v4, ok := vals[3].(*big.Int)
		if !ok {
			return nil, suppliedGas, errors.New("invalid val")
		}
		v5, ok := vals[4].(*big.Int)
		if !ok {
			return nil, suppliedGas, errors.New("invalid val")
		}
		v6, ok := vals[5].(*big.Int)
		if !ok {
			return nil, suppliedGas, errors.New("invalid val")
		}
		// valsI := []*big.Int{v1, v2, v3, v4, v5, v6}
		xys3 := big.NewInt(0)
		addend := big.NewInt(0)
		addend.Mul(v1, v2)
		xys3.Add(xys3, addend)
		addend.Mul(v3, v4)
		xys3.Add(xys3, addend)
		addend.Mul(v5, v6)
		xys3.Add(xys3, addend)
		xys3.Mul(xys3, big.NewInt(3))
		xsys := big.NewInt(0)
		xs := big.NewInt(0)
		ys := big.NewInt(0)
		xs.Add(v1, v3)
		xs.Add(xs, v5)
		ys.Add(v2, v4)
		ys.Add(ys, v6)
		xsys.Mul(xs, ys)
		num := big.NewInt(0)
		num.Sub(xys3, xsys)
		denom := big.NewInt(0)
		x2s3 := big.NewInt(0)
		addend.Mul(v1, v1)
		x2s3.Add(x2s3, addend)
		addend.Mul(v3, v3)
		x2s3.Add(x2s3, addend)
		addend.Mul(v5, v5)
		x2s3.Add(x2s3, addend)
		x2s3.Mul(x2s3, big.NewInt(3))
		xs2 := big.NewInt(0)
		xs2.Mul(xs, xs)
		denom.Sub(x2s3, xs2)
		perc := big.NewInt(0)
		perc.Mul(num, big.NewInt(100))
		perc.Div(perc, denom)
		// fitb := 100 * (3*(v1*v2+v3*v4+v5*v6) - (v1+v3+v5)*(v2+v4+v6)) / (3*(v1*v1+v3*v3+v5*v5) - (v1+v3+v5)*(v1+v3+v5))
		// fita := 100*(v2+v4+v6)/3 - fitb*(v1+v3+v5)/3
		retb, err = MakeRetArgs().PackValues([]interface{}{perc})
		if err != nil {
			return errb[:], suppliedGas, err
		}
		return retb, suppliedGas, nil
	}
	funcGetXChainECRecover := newStatefulPrecompileFunction(fitSignature, f)
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{funcGetXChainECRecover})
	return contract
}
