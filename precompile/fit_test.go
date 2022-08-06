package precompile

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestFit(t *testing.T) {
	a, err := MakeFitArgs().PackValues([]interface{}{[]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(3), big.NewInt(2)}})
	require.NoError(t, err)
	//t.Log(hexutil.Encode(a[:]))
	res, _, err := fit(nil, common.Address{}, common.Address{}, a, 0, false)
	require.NoError(t, err)
	vls, err := MakeFitRetArgs().UnpackValues(res)
	require.NoError(t, err)
	t.Log(vls)
}
