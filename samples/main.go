package main

import (
	"fmt"
	"github.com/xiaoyueya/chain-sdk/common"
	"github.com/xiaoyueya/chain-sdk/services"
)

/**
"redis_config":{
    "Addr":"127.0.0.1:6379",
    "DB":1,
    "DialTimeout":10,
    "ReadTimeout":30,
    "WriteTimeout":30,
    "PoolSize":10,
    "PoolTimeout":30,
    "IdleTimeout":500,
    "IdleCheckFrequency":500,
    "Password":""
  },
**/
func main() {
	err := services.InitRedisService(&common.RedisConfig{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
		DialTimeout:        10,
		ReadTimeout:        30,
		WriteTimeout:       30,
		PoolSize:           10,
		PoolTimeout:        30,
		IdleTimeout:        500,
		IdleCheckFrequency: 500,
		Password:           "654123",
	})

	if err != nil {
		panic(err)
	}

	//chainInfo := &common.ChainInfo{
	//	Chain:                   "FTM",
	//	ChainID:                 "250",
	//	RPCUrls:                 []string{"https://rpc.ftm.tools/"},
	//	Symbol:                  "FTM",
	//	BlockExplorerAddressURL: "https://ftmscan.com/address/%s",
	//	BlockExplorerTxURL:      "https://ftmscan.com/tx/%s",
	//	GasCustomCoin:           "FTM",
	//}
	chainInfo := &common.ChainInfo{
		Chain:                   "FTM",
		ChainID:                 "4002",
		RPCUrls:                 []string{"https://rpc.testnet.fantom.network/"},
		Symbol:                  "FTM",
		BlockExplorerAddressURL: "https://testnet.ftmscan.com/address/%s",
		BlockExplorerTxURL:      "https://testnet.ftmscan.com/tx/%s",
		GasCustomCoin:           "FTM",
	}

	fmt.Printf("chain info=%v\n", chainInfo)
	err = services.AddOrUpdateChain(chainInfo)
	if err != nil {
		panic(err)
	}
	coinInfo := &common.CoinOrToken{
		Chain:           "FTM",
		ChainId:         "4002",
		CustomName:      "FTM",
		TokenName:       "Fantom Coin",
		Symbol:          "FTM",
		Decimals:        18,
		ContractAddress: "",
		IsErc20:         false,
	}
	err = services.AddOrUpdateCoin(coinInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("coin info=%v\n", coinInfo)
	info, err := services.GetChain(chainInfo.Chain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("info=%v\n", info)

	cInfo, err := services.GetCoinOrToken(coinInfo.CustomName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cinfo=%v\n", cInfo)

	coinInfo = &common.CoinOrToken{
		Chain:           "FTM",
		ChainId:         "250",
		CustomName:      "FRC20USDC",
		TokenName:       "USD Coin",
		Symbol:          "USDC",
		Decimals:        6,
		ContractAddress: "0x9740c9375200ace24af03d7aa488b031877735c5",
		IsErc20:         true,
	}
	err = services.AddOrUpdateCoin(coinInfo)
	if err != nil {
		panic(err)
	}

	tokens, err := services.GetAllXrc20s()
	if err != nil {
		panic(err)
	}
	fmt.Printf("tokens=%v\n", tokens)
}
