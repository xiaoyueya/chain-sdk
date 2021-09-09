package common

import "fmt"

const (
	BucketChain = "chain:%s:%s"	//chain:链简称：链id，存储链信息
	BucketCoin  = "coin:%s:%s:%s"	//coin:链名：链id:币种简称，存储币种信息
)

func BuildChainKey(chain, chainId string) string {
	return fmt.Sprintf(BucketChain,chain,chainId)
}

func BuildCoinKey(chain, chainId,coin string) string {
	return fmt.Sprintf(BucketCoin,chain,chainId,coin)
}