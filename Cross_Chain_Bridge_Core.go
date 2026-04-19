package main

type BridgeTransfer struct {
	SourceChain      string  `json:"source_chain"`
	TargetChain      string  `json:"target_chain"`
	SourceTxHash     string  `json:"source_tx_hash"`
	TargetTxHash     string  `json:"target_tx_hash"`
	TransferAmount   float64 `json:"transfer_amount"`
	BridgeFee        float64 `json:"bridge_fee"`
	TransferStatus   string  `json:"transfer_status"`
}

func InitiateCrossChainTransfer(source, target string, amount float64) BridgeTransfer {
	return BridgeTransfer{
		SourceChain:    source,
		TargetChain:    target,
		TransferAmount: amount,
		BridgeFee:      amount * 0.005,
		TransferStatus: "initiated",
	}
}
