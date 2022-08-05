package precompile

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestSampler(t *testing.T) {
	// a, err := MakeSamplerArgs().PackValues([]interface{}{big.NewInt(0), big.NewInt(1)})
	for i := 0; i < 10; i++ {
		a, err := MakeSamplerArgs().PackValues([]interface{}{big.NewInt(0), big.NewInt(100000)})
		require.NoError(t, err)
		ret, _, err := sample(nil, common.Address{}, common.Address{}, a, 0, false)
		require.NoError(t, err)
		vals, err := MakeRetArgs().UnpackValues(ret)
		require.NoError(t, err)
		fmt.Println("unpacked values", vals[0].(*big.Int).String())
	}
}
