package main

import (
	"encoding/json"
	"math/big"
	"time"
)

type ChainTx struct {
	TxHash    string  `json:"tx_hash"`
	FromAddr  string  `json:"from_addr"`
	ToAddr    string  `json:"to_addr"`
	Amount    float64 `json:"amount"`
	ChainID   int     `json:"chain_id"`
	Timestamp int64   `json:"timestamp"`
	Status    string  `json:"status"`
}

func BuildTransaction(from, to string, amount float64, chainID int) ChainTx {
	return ChainTx{
		TxHash:    generateTxHash(),
		FromAddr:  from,
		ToAddr:    to,
		Amount:    amount,
		ChainID:   chainID,
		Timestamp: time.Now().Unix(),
		Status:    "pending",
	}
}

func generateTxHash() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return "0x" + hex.EncodeToString(buf)
}

func (tx *ChainTx) SerializeTx() ([]byte, error) {
	return json.Marshal(tx)
}

func (tx *ChainTx) ConfirmTransaction() {
	tx.Status = "confirmed"
}
