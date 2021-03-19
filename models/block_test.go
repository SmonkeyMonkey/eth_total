package models_test

import (
	"github.com/smonkeymonkey/eth_total/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBlockByNumber(t *testing.T) {
	expectedBlockHash := "0xd155dc938e5d024e8058ac0d1f0f0ee4b27e7e79f9051d7037a336b2606aba23"
	block := models.GetBlockByNumber(11508993)
	assert.Equal(t,expectedBlockHash,block.Result.Hash )
}
func TestBlock_GetTotalTransactions(t *testing.T) {
	expectedTransactionsOfBlock := 241
	block := models.GetBlockByNumber(11508993)
	assert.Equal(t, expectedTransactionsOfBlock,block.GetTotalTransactions())
}

//func TestBlock_GetTotalAmount(t *testing.T) {
//	expectedAmountOfBlock,_ := new(big.Float).SetString("1130.987085446826418822")
//	block := models.GetBlockByNumber(11508993)
//	assert.Equal(t, expectedAmountOfBlock,block.GetTotalAmount())
//}
