package wallet_analyser

import (
	"example.com/data_lib"
)

func FilterTxs(transactions *data_lib.Response, addr string) []string {
	hashes := make([]string, 0)

	for _, tx := range transactions.Result {
		if tx.To == addr {
			hashes = append(hashes, tx.Hash)
		}
	}

	return hashes
}
