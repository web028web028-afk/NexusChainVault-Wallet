package main

import "encoding/hex"

type ERC20Handler struct {
	ContractAddress string
}

func (e *ERC20Handler) EncodeTransferMethod(to string, amount *big.Int) string {
	methodID := "a9059cbb"
	paddedTo := padLeft(hex.EncodeToString([]byte(to[2:])), 64)
	paddedAmount := padLeft(amount.Text(16), 64)
	return "0x" + methodID + paddedTo + paddedAmount
}

func padLeft(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}
	return s
}

func (e *ERC20Handler) ParseTransferEvent(data string) (string, *big.Int) {
	to := "0x" + data[26:66]
	amount, _ := new(big.Int).SetString(data[66:], 16)
	return to, amount
}
