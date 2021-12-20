package ethClientHelper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
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

func (ec *EthClient) InitEthClient() error {
	// ethClient, err := ethclient.Dial("https://rinkeby.infura.io/v3/1fb06dad21b2437584a5426c0dced67b")
	// // ethClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/1fb06dad21b2437584a5426c0dced67b")
	// ethClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/1fb06dad21b2437584a5426c0dced67b")
	// // ethClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/1fb06dad21b2437584a5426c0dced67b")
	ethClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/1fb06dad21b2437584a5426c0dced67b")
	if err != nil {
		return err
	}
	ec.EthClient = ethClient
	return nil
}

// if returns true, address is valid contract address. If false, address is standard ethereum account
func (ec *EthClient) CheckAddressType(_address string) (bool, error) {
	address := common.HexToAddress(_address)
	bc, err := ec.EthClient.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, err
	}

	isContract := len(bc) > 0
	return isContract, nil
}

func GetContractABI(_address string) (string, error) {
	resp, err := http.Get("https://api.etherscan.io/api?module=contract&action=getsourcecode&address=" + _address)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	var obj EtherscanRespObj
	if err := json.Unmarshal(body, &obj); err != nil {
		log.Fatalln(err)
		return "", err
	}

	if obj.Status != "1" && len(obj.Results) == 0 {
		log.Fatalln(err)
		return "", err
	}

	return obj.Results[0].ABI, nil
}
