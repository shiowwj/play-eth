package ethClientHelper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// if returns true, address is valid contract address. If false, address is standard ethereum account
func (ec *EthClient) CheckAddressType(address string) (bool, error) {
	a := common.HexToAddress(address)
	bc, err := ec.EthClient.CodeAt(context.Background(), a, nil)
	if err != nil {
		return false, err
	}

	isContract := len(bc) > 0
	return isContract, nil
}

func GetContractABI(address string) (string, error) {
	resp, err := http.Get("https://api.etherscan.io/api?module=contract&action=getsourcecode&address=" + address)
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

// ToWei decimals to wei
func ToWei(iAmt interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iAmt.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
