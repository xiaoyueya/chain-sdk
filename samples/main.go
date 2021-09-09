package main

import (
	"fmt"
	"github.com/xiaoyueya/chain-sdk/common"
)

func main() {
	chainInfo := common.ChainInfo{}
	fmt.Printf("chain info=%v\n", chainInfo)
	coinInfo := common.CoinOrToken{}
	fmt.Printf("coin info=%v\n", coinInfo)

}
