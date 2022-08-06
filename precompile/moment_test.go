package precompile

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestMoment(t *testing.T) {
	input, err := MakeMomentArgs().Pack(big.NewInt(2), []*big.Int{big.NewInt(23), big.NewInt(31), big.NewInt(23)}, []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1)})
	require.NoError(t, err)
	med, s, s2 := calcMoment(nil, common.Address{}, common.Address{}, input, 100_000, true)
	require.NoError(t, err)
	t.Log(med)
}
