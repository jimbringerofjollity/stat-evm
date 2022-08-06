package precompile

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type Distribution interface {
	Rand() float64
}

type ContractSamplerConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractSamplerConfig) Address() common.Address {
	return ContractSamplerAddress
}

func (c ContractSamplerConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractSamplerConfig) IsDisabled() bool {
	return false
}

func (c ContractSamplerConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractSamplerConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractSamplerConfig) Contract() StatefulPrecompiledContract {
	return ContractSamplerPrecompile
}

var (
	ContractSamplerPrecompile StatefulPrecompiledContract = createSamplerPrecompile(ContractSamplerAddress)
	samplerSignature                                      = CalculateFunctionSelector("sampleRandomNumber(uint256,int256,int256,uint256)")
)

func MakeSamplerArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "distributionType",
			Type: mustType("uint256"),
		},
		{
			Name: "a",
			Type: mustType("int256"),
		},
		{
			Name: "b",
			Type: mustType("int256"),
		},
		{
			Name: "numSamples",
			Type: mustType("uint256"),
		},
	}
}

func MakeSamplerReturnArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("int256[]"),
		},
	}
}

func sampleRandomNumber(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	// TODO: Figure out gas cost as a function of numSamples

	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)

	inputValues, err := MakeSamplerArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	if len(inputValues) != 4 {
		return nil, suppliedGas, errors.New("invalid vals")
	}

	distributionType, ok := inputValues[0].(*big.Int)
	intDistributionType := int(distributionType.Int64())
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}

	a, ok := inputValues[1].(*big.Int)
	floatA := float64(a.Int64())
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}

	b, ok := inputValues[2].(*big.Int)
	floatB := float64(b.Int64())
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}

	numSamples, ok := inputValues[3].(*big.Int)
	intNumSamples := int(numSamples.Int64())
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}

	// Get random seed from hash of caller address, receiver address, and block timestamp
	h := md5.New()
    io.WriteString(h, callerAddr.String() + addr.String() + evm.GetBlockContext().Timestamp().String())
	randomSeed := rand.NewSource(binary.BigEndian.Uint64(h.Sum(nil)))

	// Declare distribution variable to use later
	var dist Distribution

	switch intDistributionType {
	case 0:
		// Uniform distribution
		dist = distuv.Uniform{
			Min: floatA, // Mean of the normal distribution
			Max: floatB, // Standard deviation of the normal distribution
			Src: randomSeed,
		}
	case 1:
		// Normal distribution
		dist = distuv.Normal{
			Mu:    floatA, // Mean of the normal distribution
			Sigma: floatB, // Standard deviation of the normal distribution
			Src:   randomSeed,
		}
	case 2:
		// Binomial distribution
		dist = distuv.Binomial{
			N:   floatA,
			P:   floatB,
			Src: randomSeed,
		}
	case 3:
		// Beta distribution
		dist = distuv.Beta{
			Alpha: floatA,
			Beta:  floatB,
			Src:   randomSeed,
		}
	default:
		// Uniform distribution
		dist = distuv.Uniform{
			Min: floatA, // Mean of the normal distribution
			Max: floatB, // Standard deviation of the normal distribution
			Src: randomSeed,
		}
	}

	// Sample numSamples from distribution
	samples := make([]float64, intNumSamples)
	bigValues := make([]*big.Int, intNumSamples)
	for i := 0; i < intNumSamples; i++ {
		samples[i] = dist.Rand()
		bigValue := new(big.Float).SetFloat64(samples[i])
		f, _ := bigValue.Int64()
		bigValues[i] = new(big.Int).SetInt64(f)
	}

	ret, err = MakeSamplerReturnArgs().PackValues([]interface{}{bigValues})
	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}

func createSamplerPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(samplerSignature, sampleRandomNumber)})
}
