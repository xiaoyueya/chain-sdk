package common

type ChainInfo struct {
	Chain                   string   `json:"chain"`
	RPCUrls                 []string `json:"rpc_urls"`
	ChainID                 string   `json:"chain_id"`
	Symbol                  string   `json:"symbol"`
	BlockExplorerTxURL      string   `json:"block_explorer_tx_url"`
	BlockExplorerAddressURL string   `json:"block_explorer_address_url"`
}

type CoinOrToken struct {
	Chain           string `json:"chain"`
	ChainId         string `json:"chain_id"`
	CustomName      string `json:"custom_name"`
	TokenName       string `json:"token_name"`
	Symbol          string `json:"symbol"`
	Decimals        uint   `json:"decimals"`
	ContractAddress string `json:"contract_address"`
	IsXXX20         bool   `json:"is_xxx_20"`
}

type RedisConfig struct {
	Addr               string
	DB                 int
	DialTimeout        int32
	ReadTimeout        int32
	WriteTimeout       int32
	PoolSize           int
	PoolTimeout        int32
	IdleTimeout        int32
	IdleCheckFrequency int32
	Password           string
}
