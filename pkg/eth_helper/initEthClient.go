package ethClientHelper

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	EthClient *ethclient.Client
}

type EtherscanRespObj struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Results []Result `json:"result"`
}

type Result struct {
	SourceCode           string `json:"SourceCode"`
	ABI                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	ConstructorArguments string `json:"ConstructorArguments"`
}

func (ec *EthClient) InitEthClient() (err error) {
	// for rinkeby
	ethClient, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/d58d510829ed4da88495c6b227c1b0fa")

	// for mainnet
	// ethClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/1fb06dad21b2437584a5426c0dced67b")

	if err != nil {
		return err
	}
	ec.EthClient = ethClient
	return nil
}
