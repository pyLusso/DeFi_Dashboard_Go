package data_lib

type Transaction struct {
	BlockNumber       int
	TimeStamp         int
	Hash              string
	Nonce             int
	BlockHash         string
	TransactionIndex  int
	From              string
	To                string
	Value             int
	Gas               int
	GasPrice          int
	isError           int
	TxReceiptStatus   int
	Input             string
	ContractAddess    string
	CumulativeGasUsed int
	GasUsed           int
	Confirmations     int
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}
