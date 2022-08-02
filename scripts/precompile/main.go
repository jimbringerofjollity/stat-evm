package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ava-labs/subnet-evm/ethclient"
)

func confirm(ec ethclient.Client, txHash common.Hash) {
	for {
		_, pending, err2 := ec.TransactionByHash(context.Background(), txHash)
		fmt.Println(pending, err2)
		if !pending {
			receipt, err3 := ec.TransactionReceipt(context.Background(), txHash)
			fmt.Println(err3, receipt.GasUsed)
			break
		}
		time.Sleep(time.Second)
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ec, err := ethclient.Dial("http://127.0.0.1:12205/ext/bc/2hRGpoeHaEEQ7F9gJSzwkdCcmygRkS4ruWkR8i3w8LsFua4iUG/rpc")
	panicErr(err)
	precompile := common.HexToAddress("0x0300000000000000000000000000000000000001")
	//tx := types.NewTransaction(
	//	1, precompile, nil, 500_000, nil, hexutil.MustDecode("0x1ff7dabc"))
	//types.NewTx(types.NewTx())
	err = ec.SendTransaction(context.Background(), tx)
	panicErr(err)
	confirm(ec, tx.Hash())
	//b, err := ec.ChainID(context.Background())
	//panicErr(err)
	//key, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	//panicErr(err)
	//user, err := bind.NewKeyedTransactorWithChainID(key, b)
	//panicErr(err)
	//addr, deployTx, testContract, err := DeployTest(user, ec)
	//panicErr(err)
	//fmt.Println("address", addr)
	//confirm(ec, deployTx.Hash())
	////testContract, _ := NewTest(common.HexToAddress("0x789a5FDac2b37FCD290fb2924382297A6AE65860"), ec)
	//user.GasLimit = 500_000
	//tx, err := testContract.TestMe(user)
	//panicErr(err)
	//confirm(ec, tx.Hash())
	//l, err := testContract.Last(nil)
	//panicErr(err)
	//fmt.Println(hexutil.Encode(l[:]), err)
}
