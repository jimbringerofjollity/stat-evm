package precompile

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	ContractFitPrecompile StatefulPrecompiledContract = createFitPrecompile(ContractFitAddress)
	fitSignature                                      = crypto.Keccak256([]byte("getFit(uint256[])"))[:4] // Should be fit(uint256[][],uint256[][],string)
)

func MakeFitArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "vals",
			Type: mustType("uint256[]"),
		},
	}
}

func MakeFitRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func fit(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (retb []byte, remainingGas uint64, err error) {
	fmt.Println("input", hexutil.Encode(input))
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	valsI, err := MakeFitArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}

	vals, ok := valsI[0].([]*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid type")
	}
	if len(vals)%2 != 0 {
		return nil, suppliedGas, errors.New("invalid parity of len(vals)")
	}

	xys3 := big.NewInt(0)
	for j := 0; j < len(vals); j += 2 {
		addend := big.NewInt(0)
		addend.Mul(vals[j], vals[j+1])
		xys3.Add(xys3, addend)
	}
	xys3.Mul(xys3, big.NewInt(3))
	xsys := big.NewInt(0)
	xs := big.NewInt(0)
	ys := big.NewInt(0)
	for j := 0; j < len(vals); j += 2 {
		xs.Add(xs, vals[j])
		ys.Add(ys, vals[j+1])
	}
	xsys.Mul(xs, ys)
	num := big.NewInt(0)
	num.Sub(xys3, xsys)
	denom := big.NewInt(0)
	x2s3 := big.NewInt(0)
	for j := 0; j < len(vals); j += 2 {
		addend := big.NewInt(0)
		addend.Mul(vals[j], vals[j])
		x2s3.Add(x2s3, addend)
	}
	x2s3.Mul(x2s3, big.NewInt(3))
	xs2 := big.NewInt(0)
	xs2.Mul(xs, xs)
	denom.Sub(x2s3, xs2)
	perc := big.NewInt(0)
	perc.Mul(num, big.NewInt(100))
	perc.Div(perc, denom)
	// fitb := 100 * (3*(v1*v2+v3*v4+v5*v6) - (v1+v3+v5)*(v2+v4+v6)) / (3*(v1*v1+v3*v3+v5*v5) - (v1+v3+v5)*(v1+v3+v5))
	// fita := 100*(v2+v4+v6)/3 - fitb*(v1+v3+v5)/3
	retb, err = MakeFitRetArgs().PackValues([]interface{}{perc})
	if err != nil {
		return nil, suppliedGas, err
	}

	return retb, suppliedGas, nil
}

func createFitPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return the contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(fitSignature, fit)})
}
