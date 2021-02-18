package main

import (
	"fmt"

	"example.com/data_lib"
	"example.com/json_handler"
	"example.com/wallet_analyser"
)

func main() {
	WalletAddr := "0x3237a1CD9fB1d6c7f53E02B3411754D949bFCAC7"
	api_call := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=0&sort=asc&apikey=XRFWK1IDBR545CXNJ6NQSYAVINUQB7IDV1", WalletAddr)

	data_json := new(data_lib.Response)
	json_handler.GetJson(api_call, data_json)
	fmt.Printf("data_json 2 = %T %s\n", data_json.Result[0].Hash, data_json.Result[0].Hash)

	SushiSwapRouterAddr := "0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f"
	sushiswap_txs := wallet_analyser.FilterTxs(data_json, SushiSwapRouterAddr)

	fmt.Print(sushiswap_txs)

	// client, err := jsonrpc.NewClient("https://mainnet.infura.io")
	// // client, err := jsonrpc.NewClient("https://example.com")
	// fmt.Print(client)
	// fmt.Print(err)
	// if err != nil {
	// 	panic(err)
	// }

	// number, err := client.Eth().BlockNumber()
	// if err != nil {
	// 	panic(err)
	// }

	// header, err := client.Eth().GetBlockByNumber(web3.BlockNumber(number), true)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(header)
}
