package transactions

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miguelmota/go-ethutil"
)

type BlockDetail struct {
	BlockNumber       int64         `json:"blockNumber"`
	Timestamp         uint64        `json:"timestamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactionsCount"`
	Transactions      []Transaction `json:"transactions"`
}

// Transaction data structure
type Transaction struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
	From     string `json:"from"`
}

func (blk BlockDetail) GetTransactionDetails(block *types.Block, chainID *big.Int) BlockDetail {

	blk.BlockNumber = block.Number().Int64()
	blk.Timestamp = block.Time()
	blk.Difficulty = block.Difficulty().Uint64()
	blk.Hash = block.Hash().Hex()
	blk.TransactionsCount = len(block.Transactions())
	if blk.TransactionsCount == 0 {
		// add logger for no transaction count
		return blk
	}

	txns := []Transaction{}
	for _, _txn := range block.Transactions() {
		txn := Transaction{}
		txnObj := txn.getTransactionDetail(_txn, chainID)
		txns = append(txns, txnObj)
	}

	blk.Transactions = txns
	// add loggere here to check blk details + blk's transactions
	return blk
}

func (txn Transaction) getTransactionDetail(_txn *types.Transaction, chainID *big.Int) Transaction {
	// fBalance := new(big.Float)
	// fBalance.SetString(_txn.Value().String())
	// ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	ethValue := ethutil.ToDecimal(_txn.Value(), 18)
	txn.Hash = _txn.Hash().Hex()
	if _txn.To() != nil {
		txn.To = _txn.To().Hex()
	}
	// txn.Value = ethValue.String()
	txn.Value = ethValue.String()
	txn.Gas = _txn.Gas()
	txn.GasPrice = _txn.GasPrice().Uint64()
	txn.Nonce = _txn.Nonce()

	if msg, err := _txn.AsMessage(types.NewEIP155Signer(chainID), big.NewInt(1)); err == nil {
		txn.From = msg.From().Hex()
	}
	// should add logger for txn details here?
	return txn
}
