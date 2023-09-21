package withGETH

import "math/big"

type Transaction struct {
	Address              string        `bson:"address"`
	Type                 uint8         `bson:"type"`
	ChainId              *big.Int      `bson:"chain_id"`
	Nonce                string        `bson:"nonce"`
	To                   string        `bson:"to"`
	Gas                  string        `bson:"gas"`
	GasPrice             interface{}   `bson:"gas_price"`
	MaxPriorityFeePerGas string        `bson:"max_priority_fee_per_gas"`
	MaxFeePerGas         string        `bson:"max_fee_per_gas"`
	Value                string        `bson:"value"`
	Input                string        `bson:"input"`
	AccessList           []interface{} `bson:"access_list"`
	V                    string        `bson:"v"`
	R                    string        `bson:"r"`
	S                    string        `bson:"s"`
	YParity              string        `bson:"y_parity"`
	Hash                 string        `bson:"hash"`
}
