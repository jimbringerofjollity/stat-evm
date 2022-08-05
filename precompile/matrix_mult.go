package precompile

import "C"
import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gonum.org/v1/gonum/mat"
)

type ContractMatrixMultConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractMatrixMultConfig) Address() common.Address {
	return ContractMatrixMultAddress
}

func (c ContractMatrixMultConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractMatrixMultConfig) IsDisabled() bool {
	return false
}

func (c ContractMatrixMultConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractMatrixMultConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractMatrixMultConfig) Contract() StatefulPrecompiledContract {
	return ContractMatrixMultPrecompile
}

var (
	ContractMatrixMultPrecompile StatefulPrecompiledContract = createMatrixMultPrecompile(ContractMatrixMultAddress)
	matrixSignature                                          = crypto.Keccak256([]byte("matrixMultiply(int256[][],int256[][])"))[:4]
)

func mustMatrixMultType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeMatrixMultArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "a",
			Type: mustMatrixMultType("int256[][]"),
		},
		{
			Name: "b",
			Type: mustMatrixMultType("int256[][]"),
		},
	}
}

func MakeMatrixMultRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "c",
			Type: mustType("int256[][]"),
		},
	}
}

func buildMatrix(a [][]*big.Int) mat.Matrix {
	var data []float64
	for row := 0; row < len(a); row++ {
		for col := 0; col < len(a[row]); col++ {
			data = append(data, float64(a[row][col].Int64()))
		}
	}
	return mat.NewDense(len(a), len(a[0]), data)
}

func matrixMult(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	vals, err := MakeMatrixMultArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	if len(vals) != 2 {
		return nil, suppliedGas, errors.New("expected two vals")
	}
	a, ok := vals[0].([][]*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	b, ok := vals[1].([][]*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	var c mat.Dense
	c.Mul(buildMatrix(a), buildMatrix(b))
	nr, _ := c.Dims()
	var res [][]*big.Int
	for row := 0; row < nr; row++ {
		rowVals := c.RawRowView(row)
		var solRow []*big.Int
		for _, rowVal := range rowVals {
			i, _ := big.NewFloat(rowVal).Int64()
			solRow = append(solRow, big.NewInt(i))
		}
		res = append(res, solRow)
	}
	ret, err = MakeMatrixMultRetArgs().PackValues([]interface{}{res})
	if err != nil {
		return nil, suppliedGas, err
	}
	return ret, suppliedGas, nil
}

func createMatrixMultPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(matrixSignature, matrixMult)})
}
