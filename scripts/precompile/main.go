package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func confirm(ec *ethclient.Client, txHash common.Hash) {
	for {
		_, pending, err2 := ec.TransactionByHash(context.Background(), txHash)
		fmt.Println("Pending:", pending)
		fmt.Println("Error:", err2)
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

func start_events(ec *ethclient.Client, addr common.Address) {
	ec, err := ethclient.Dial("ws://127.0.0.1:27044/ext/bc/t5jkcEANP8EXobp3MRGeDiBZpuQ4Rp8359Vd5vVaagdgwdW4u/ws")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	logs := make(chan types.Log)
	sub, err := ec.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			fmt.Println()
		}
	}
}

func main() {
	ec, err := ethclient.Dial("http://127.0.0.1:11987/ext/bc/Vf9noCc7xT5viKRcvhhxK9JQ2LzVn1FLWU4uzGgwso99cybXu/rpc")
	panicErr(err)

	b, err := ec.ChainID(context.Background())
	panicErr(err)

	key, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	panicErr(err)

	user, err := bind.NewKeyedTransactorWithChainID(key, b)
	panicErr(err)

	addr, deployTx, testContract, err := DeployMain(user, ec)
	panicErr(err)

	fmt.Println("address", addr)

	confirm(ec, deployTx.Hash())

	user.GasLimit = 500_000
	tx, err := testContract.TestMedian(user, []*big.Int{big.NewInt(5), big.NewInt(10), big.NewInt(3), big.NewInt(7), big.NewInt(12)})
	panicErr(err)
	confirm(ec, tx.Hash())
	fmt.Println("Tx hash (median):", tx.Hash())
	l, err := testContract.Med(nil)
	panicErr(err)
	fmt.Println("median", l, err)

	sampler_tx, err := testContract.TestSampler(user, big.NewInt(1), big.NewInt(1000000000000000000))
	panicErr(err)
	confirm(ec, sampler_tx.Hash())
	fmt.Println("Tx hash (sampler):", sampler_tx.Hash())

	sampler_res, err := testContract.Sample(nil)
	panicErr(err)
	fmt.Println("sample result", sampler_res, err)

	var a [][]*big.Int
	for i := 0; i < 1; i++ {
		a = append(a, []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)})
	}
	var b2 [][]*big.Int
	for i := 0; i < 3; i++ {
		b2 = append(b2, []*big.Int{big.NewInt(int64(i + 1))})
	}
	matrix_tx, err := testContract.TestMatrixMult(user, a, b2)
	panicErr(err)
	confirm(ec, matrix_tx.Hash())
	fmt.Println("Tx hash (matrix):", matrix_tx.Hash())
	vals, err := testContract.GetMatrixMulti(nil)
	panicErr(err)
	fmt.Println(vals)
}
