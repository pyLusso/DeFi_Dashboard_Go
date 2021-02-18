package main

import (
	"context"
	"fmt"

	"example.com/data_lib"
	"example.com/json_handler"
	"example.com/wallet_analyser"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Specify address and format api call
	WalletAddr := common.HexToAddress("0x3237a1CD9fB1d6c7f53E02B3411754D949bFCAC7")

	api_call := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=0&sort=asc&apikey=XRFWK1IDBR545CXNJ6NQSYAVINUQB7IDV1", WalletAddr)

	// Initialise empty response data structure
	data_json := new(data_lib.Response)

	// Map response onto data interface
	json_handler.GetJson(api_call, data_json)

	// Specify sushiswap router address
	SushiSwapRouterAddr := "0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f"

	// Filter for transactions that go to sushiswap
	sushiswap_txs := wallet_analyser.FilterTxs(data_json, SushiSwapRouterAddr)

	fmt.Print("\n These txs went to sushiswap: ", sushiswap_txs, "\n")

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/f9a81520189642c89e0e2163ede73662")
	// client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	fmt.Println("\n We have a connection!\n")
	_ = client

	for _, tx := range sushiswap_txs {
		txHash := common.HexToHash(tx)
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			panic(err)
		}
		fmt.Print("\n\n", receipt.Logs, "\n\n")

	}

}
