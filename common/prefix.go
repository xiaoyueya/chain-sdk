package common

import "fmt"

const (
	BucketChain = "chains:%s" //chain:链简称：链id，存储链信息
	BucketCoin  = "coins:%s"  //coin:链名：链id:币种简称，存储币种信息
)

func BuildChainKey(chain string) string {
	return fmt.Sprintf(BucketChain, chain)
}

func BuildCoinKey(coin string) string {
	return fmt.Sprintf(BucketCoin, coin)
}
