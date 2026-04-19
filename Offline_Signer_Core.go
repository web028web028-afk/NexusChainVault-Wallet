package main

import "encoding/json"

type OfflineSignRequest struct {
	TxData    string `json:"tx_data"`
	ChainID   int    `json:"chain_id"`
	Timestamp int64  `json:"timestamp"`
}

func CreateOfflineSignRequest(txData string, chainID int) OfflineSignRequest {
	return OfflineSignRequest{
		TxData:    txData,
		ChainID:   chainID,
		Timestamp: time.Now().Unix(),
	}
}

func (o *OfflineSignRequest) ExportForSigning() []byte {
	data, _ := json.Marshal(o)
	return data
}

func VerifyOfflineRequest(req OfflineSignRequest) bool {
	return req.Timestamp > time.Now().Unix()-3600
}
