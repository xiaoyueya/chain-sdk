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

	chainInfo := &common.ChainInfo{
		Chain:                   "FTM",
		ChainID:                 "250",
		RPCUrls:                 []string{"https://rpc.ftm.tools/"},
		Symbol:                  "FTM",
		BlockExplorerAddressURL: "https://ftmscan.com/address/%s",
		BlockExplorerTxURL:      "https://ftmscan.com/tx/%s",
	}
	fmt.Printf("chain info=%v\n", chainInfo)
	err = services.AddOrUpdateChain(chainInfo)
	if err != nil {
		panic(err)
	}
	coinInfo := &common.CoinOrToken{
		Chain:           "FTM",
		ChainId:         "250",
		CustomName:      "FRC20USDC",
		TokenName:       "USD coin",
		Symbol:          "USDC",
		Decimals:        6,
		ContractAddress: "0x04068da6c83afcfa0e13ba15a6696662335d5b75",
		IsErc20:         true,
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
}
