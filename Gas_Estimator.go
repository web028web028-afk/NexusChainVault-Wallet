package main

import "math/big"

type GasService struct {
	rpc *RPCClient
}

func NewGasService(rpc *RPCClient) *GasService {
	return &GasService{rpc: rpc}
}

func (g *GasService) EstimateGasLimit() (*big.Int, error) {
	res, err := g.rpc.CallRPC("eth_estimateGas", []interface{}{})
	if err != nil {
		return nil, err
	}
	hexGas := res["result"].(string)
	gas := new(big.Int)
	gas.SetString(hexGas[2:], 16)
	return gas, nil
}

func (g *GasService) GetGasPrice() (*big.Float, error) {
	res, err := g.rpc.CallRPC("eth_gasPrice", []interface{}{})
	if err != nil {
		return nil, err
	}
	hexPrice := res["result"].(string)
	wei := new(big.Int)
	wei.SetString(hexPrice[2:], 16)
	gwei := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e9))
	return gwei, nil
}
