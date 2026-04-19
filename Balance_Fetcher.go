package main

import "math/big"

type BalanceService struct {
	rpc *RPCClient
}

func NewBalanceService(rpc *RPCClient) *BalanceService {
	return &BalanceService{rpc: rpc}
}

func (b *BalanceService) GetNativeBalance(address string) (*big.Float, error) {
	res, err := b.rpc.CallRPC("eth_getBalance", []interface{}{address, "latest"})
	if err != nil {
		return nil, err
	}
	hexBal := res["result"].(string)
	wei := new(big.Int)
	wei.SetString(hexBal[2:], 16)
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return eth, nil
}

func (b *BalanceService) GetTokenBalance(contract, address string) (*big.Float, error) {
	return big.NewFloat(0), nil
}
