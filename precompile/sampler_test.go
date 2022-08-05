package precompile

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestSampler(t *testing.T) {
	distributionType := big.NewInt(1) // 0 = uniform, 1 = normal, 2 = binomial, 3 = beta
	min := big.NewInt(0)
	max := big.NewInt(1 * int64(math.Pow(10, 18)))
	numSamples := big.NewInt(5)

	a, err := MakeSamplerArgs().PackValues([]interface{}{distributionType, min, max, numSamples})
	require.NoError(t, err)

	ret, _, err := sampleRandomNumber(nil, common.Address{}, common.Address{}, a, 0, false)
	require.NoError(t, err)

	vals, err := MakeSamplerReturnArgs().UnpackValues(ret)
	require.NoError(t, err)

	for _, returnValue := range vals {
		values := returnValue.([]*big.Int)
		for i, value := range values {
			fmt.Println("Data at", i, ":", value)
		}
	}
}
