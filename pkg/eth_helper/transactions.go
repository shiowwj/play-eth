package ethClientHelper

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

func (ec *EthClient) TransferEthToPrivateWallet(privateKey, toAddress, valueInEth string) (string, error) {

	ethClient := ec.EthClient

	pKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	pubKey := pKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Error("error casting public key to ECDSA")
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	// get the nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	//set value of txn, convert eth to wei
	value := ToWei(valueInEth, 18)
	// set gas limit ** can put a config to gaslimit
	gasLimit := uint64(21000)
	// get current gasprice average
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAdd := common.HexToAddress(toAddress)

	tx := types.NewTransaction(nonce, toAdd, value, gasLimit, gasPrice, nil)

	chainId, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), pKey)
	if err != nil {
		return "", err
	}

	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
