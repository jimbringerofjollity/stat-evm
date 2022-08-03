package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func confirm(ec *ethclient.Client, txHash common.Hash) {
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
	ec, err := ethclient.Dial("http://127.0.0.1:15257/ext/bc/DQFh4QJkpgKRbVi9x8VeQUwNnWbGbXpLizZLid71hNJB8MwHw/rpc")
	panicErr(err)
	b, err := ec.ChainID(context.Background())
	panicErr(err)
	key, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	panicErr(err)
	user, err := bind.NewKeyedTransactorWithChainID(key, b)
	panicErr(err)
	addr, deployTx, testContract, err := DeployTest(user, ec)
	panicErr(err)
	fmt.Println("address", addr)
	confirm(ec, deployTx.Hash())
	//testContract, _ := NewTest(common.HexToAddress("0x789a5FDac2b37FCD290fb2924382297A6AE65860"), ec)
	user.GasLimit = 500_000
	tx, err := testContract.TestMe(user, big.NewInt(1), big.NewInt(1), big.NewInt(3), big.NewInt(2), big.NewInt(2), big.NewInt(3))
	panicErr(err)
	confirm(ec, tx.Hash())
	l, err := testContract.Last(nil)
	panicErr(err)
	fmt.Println(l, err)
}
