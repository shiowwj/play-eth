package main

import (
	"log"
	"math/big"
	"os"

	"github.com/joho/godotenv"
	ethClientHelper "github.com/shiowwj/server_m1/pkg/eth_helper"
)

// var mainnetChainId = big.NewInt(1) // for mainnet
var rinkebyChainId = big.NewInt(4) // for mainnet

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	//
	ec := ethClientHelper.EthClient{}
	err = ec.InitEthClient()
	if err != nil {
		log.Fatal(err)
	}

	txnHash, err := ec.TransferEthToPrivateWallet(os.Getenv("TEST_ACC_TWO_RINKEBY_PKEY"), "0x8BA96639822Dd6149f4494301f83590d4E667CDC", "0.001")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tx Sent: ", txnHash)

	// for main net
	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, mainnetChainId)
	// for testnet
	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, rinkebyChainId)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)
	// auth.GasLimit = uint64(300000)
	// auth.GasPrice = gasPrice

	// this part starts the contract interaction portion
	// address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	// instance, err := storeContract.NewStoreContract(address, ethClient)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// key := [32]byte{}
	// value := [32]byte{}
	// copy(key[:], []byte("foo"))
	// copy(value[:], []byte("bar"))

	// tx, err := instance.SetItem(auth, key, value)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("tx sent: %s", tx.Hash().Hex())

	// result, err := instance.Items(&bind.CallOpts{}, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(string(result[:]))
}
