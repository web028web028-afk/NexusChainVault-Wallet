package main

type ChainConfig struct {
	ChainName         string `json:"chain_name"`
	ChainID           int    `json:"chain_id"`
	RPCEndpoint       string `json:"rpc_endpoint"`
	ExplorerURL       string `json:"explorer_url"`
	Symbol            string `json:"symbol"`
	IsTestnet         bool   `json:"is_testnet"`
	DefaultGasLimit   int64  `json:"default_gas_limit"`
}

func GetEthereumMainnetConfig() ChainConfig {
	return ChainConfig{
		ChainName:   "Ethereum",
		ChainID:     1,
		RPCEndpoint: "https://mainnet.infura.io/v3/your-key",
		ExplorerURL: "https://etherscan.io",
		Symbol:      "ETH",
		IsTestnet:   false,
	}
}

func GetBSCConfig() ChainConfig {
	return ChainConfig{
		ChainName:   "BNB Smart Chain",
		ChainID:     56,
		RPCEndpoint: "https://bsc-dataseed.binance.org/",
		ExplorerURL: "https://bscscan.com",
		Symbol:      "BNB",
		IsTestnet:   false,
	}
}
