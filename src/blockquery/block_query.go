package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethClientHelper "github.com/shiowwj/server_m1/pkg/eth_helper"
	"github.com/shiowwj/server_m1/pkg/output"
	"github.com/shiowwj/server_m1/pkg/transactions"
)

func main() {

	ec := ethClientHelper.EthClient{}
	err := ec.InitEthClient()
	if err != nil {
		log.Fatal(err)
	}

	ethClient := ec.EthClient
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	getLatestBlocks(*ethClient, chainID)
}

func getLatestBlocks(ethClient ethclient.Client, chainID *big.Int) {
	fmt.Println("get latest block")
	headers := make(chan *types.Header)
	sub, err := ethClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		// add logger for error
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			// add logger for error
			log.Fatal(err)
		case header := <-headers:
			go getBlockTransactions(ethClient, header.Number, chainID)
		}
	}
}

func getBlockTransactions(ethClient ethclient.Client, blockId *big.Int, chainID *big.Int) {
	// startTime := time.Now()
	fmt.Println("processBlock", blockId)
	block, err := ethClient.BlockByNumber(context.Background(), blockId)
	if err != nil {
		// add logger for error
		log.Fatal(err)
	}

	blockDetail := transactions.BlockDetail{}

	blkDetail := blockDetail.GetTransactionDetails(block, chainID)

	b, err := json.Marshal(blkDetail)
	if err != nil {
		log.Fatal(err)
	}

	cw := output.CsvOutput{}

	cw.BodyBytes = b
	cw.InitCSV(output.TXN)
	// add timer for ending here
	return
}
