package precompile

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestMedian(t *testing.T) {
	a, err := MakeFitArgs().PackValues([]interface{}{big.NewInt(10), big.NewInt(2), big.NewInt(3)})
	require.NoError(t, err)
	t.Log(hexutil.Encode(a[:]))
	vls, err := MakeFitArgs().UnpackValues(a)
	require.NoError(t, err)
	t.Log(vls)
}
