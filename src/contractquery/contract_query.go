package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	ethClientHelper "github.com/shiowwj/server_m1/pkg/eth_helper"
)

func main() {
	ec := ethClientHelper.EthClient{}
	err := ec.InitEthClient()
	if err != nil {
		log.Fatal(err)
	}

	ethClient := ec.EthClient
	cAddress := common.HexToAddress("0x999e88075692bCeE3dBC07e7E64cD32f39A1D3ab")
	bc, err := ethClient.CodeAt(context.Background(), cAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bc) > 0
	fmt.Printf("is contract: %v\n", isContract)
}
