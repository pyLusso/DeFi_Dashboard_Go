module main

go 1.15

replace example.com/data_lib => ../data_lib

replace example.com/json_handler => ../json_handler

replace example.com/wallet_analyser => ../wallet_analyser

require (
	example.com/data_lib v0.0.0-00010101000000-000000000000
	example.com/json_handler v0.0.0-00010101000000-000000000000
	example.com/wallet_analyser v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.9.25
	github.com/ybbus/jsonrpc/v2 v2.1.6
)
