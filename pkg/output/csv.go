package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shiowwj/server_m1/pkg/transactions"
)

type CsvOutput struct {
	CsvWriter *csv.Writer
	File      *os.File
	Headers   []string
	BodyCsv   [][]string
	BodyBytes []byte
}

// add a return to know if csv output is successful
func (_csv *CsvOutput) InitCSV(_type string) {
	path, _ := os.Getwd()
	exPath := filepath.Dir(path)
	p := strings.ReplaceAll(exPath, "src", "")

	if _type == TXN {
		blkDetail := transactions.BlockDetail{}
		json.Unmarshal([]byte(_csv.BodyBytes), &blkDetail)
		// fmt.Printf("blkDetail InitCSV %+v\n", blkDetail)
		csvFile, err := os.Create(p + "/txn_details/test_" + strconv.Itoa(int(blkDetail.BlockNumber)) + ".csv")
		if err != nil {
			// add logger for error
			log.Fatalf("failed to create file: %s", err)
		}
		_csv.File = csvFile

		_csv.getTransactionHeaders()
		_csv.getTransactionBodyCsv(blkDetail)

		_csv.CsvWriter = csv.NewWriter(_csv.File)
		_csv.CsvWriter.Write(_csv.Headers)

		for _, a := range _csv.BodyCsv {
			err := _csv.CsvWriter.Write(a)
			if err != nil {
				// add logger for error
				fmt.Println("ERROR at CSVWRITER", err)
			}
		}
		_csv.CsvWriter.Flush()
		_csv.File.Close()
	}

	// able to add other types of csv output logic here
}

func (csv *CsvOutput) getTransactionHeaders() {
	headerRow := []string{
		"Txn Hash", "Value", "To", "From", "Gas", "Gas Price", "Nonce",
	}
	csv.Headers = headerRow
}

func (csv *CsvOutput) getTransactionBodyCsv(blk transactions.BlockDetail) {
	var txnDetails [][]string
	for _, o := range blk.Transactions {
		var txnDetail []string
		txnDetail = append(txnDetail, o.Hash)
		txnDetail = append(txnDetail, o.Value)
		txnDetail = append(txnDetail, o.To)
		txnDetail = append(txnDetail, o.From)
		txnDetail = append(txnDetail, strconv.Itoa(int(o.Gas)))
		txnDetail = append(txnDetail, strconv.Itoa(int(o.GasPrice)))
		txnDetail = append(txnDetail, strconv.Itoa(int(o.Nonce)))
		txnDetails = append(txnDetails, txnDetail)
	}
	csv.BodyCsv = txnDetails
}
