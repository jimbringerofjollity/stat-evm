package precompile

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatrixMult(t *testing.T) {
	var a [][]*big.Int
	for i := 0; i < 1; i++ { // 1 row
		a = append(a, []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)})
	}
	var b [][]*big.Int
	for i := 0; i < 3; i++ { // 1 row
		b = append(b, []*big.Int{big.NewInt(int64(i + 1))})
	}
	input, err := MakeMatrixMultArgs().PackValues([]interface{}{a, b})
	require.NoError(t, err)
	ret, _, err := matrixMult(nil, common.Address{}, common.Address{}, input, 0, false)
	require.NoError(t, err)
	v, err := MakeMatrixMultRetArgs().UnpackValues(ret)
	require.NoError(t, err)
	res := v[0].([][]*big.Int)
	assert.Equal(t, res[0][0].String(), "14")
}
