package models

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)


type Block struct {
	Result struct {
		Hash string `json:"hash"`
		Transactions []Transaction `json:"transactions"`
	}
}

type Transaction struct{
	Value string `json:"value"`
}

func GetBlockByNumber(blockNumber int) *Block{
	req := CustomRequest{}
	req.New(blockNumber)
	uri := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%s&boolean=true&apikey=%s",req.Tag,req.APIKey)
	res,_ := http.Get(uri)
	b := &Block{}
	json.NewDecoder(res.Body).Decode(&b)
	return b
}

func (b *Block) GetTotalTransactions() int {
	return len(b.Result.Transactions)
}

func(b *Block) GetTotalAmount() *big.Float{
	amountWei := new(big.Int)
	for _,v := range b.Result.Transactions {
		amountWei.Add(amountWei,b.hexToBigInt(v.Value))
	}
	return b.weiToEther(amountWei)
}

func(b *Block) hexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)

	return n
}

func(b *Block) weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
}

