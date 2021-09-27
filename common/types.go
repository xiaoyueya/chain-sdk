package common

type ChainInfo struct {
	Chain                   string   `json:"chain"`
	RPCUrls                 []string `json:"rpc_urls"`
	ChainID                 string   `json:"chain_id"`
	Symbol                  string   `json:"symbol"`
	BlockExplorerTxURL      string   `json:"block_explorer_tx_url"`
	BlockExplorerAddressURL string   `json:"block_explorer_address_url"`
	GasCustomCoin           string   `json:"gas_custom_coin"`
}

type CoinOrToken struct {
	Chain           string `json:"chain"`
	ChainId         string `json:"chain_id"`
	CustomName      string `json:"custom_name"`
	TokenName       string `json:"token_name"`
	Symbol          string `json:"symbol"`
	Decimals        uint   `json:"decimals"`
	ContractAddress string `json:"contract_address"`
	IsErc20         bool   `json:"is_erc_20"`
}

func (cot *CoinOrToken) GetContractSetKey() string {
	return BuildChainContractsKey(cot.Chain)
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

type LogConfig struct {
	Module                       string `json:"module"`
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}
