package precompile

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestMedian(t *testing.T) {
	input, err := MakeArgs().Pack([]*big.Int{big.NewInt(1), big.NewInt(4), big.NewInt(2)})
	require.NoError(t, err)
	med, _, err := getMedian(nil, common.Address{}, common.Address{}, input, 100_000, true)
	require.NoError(t, err)
	t.Log(med)
}
