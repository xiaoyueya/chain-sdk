package services

import (
	"encoding/json"
	"errors"
	"github.com/xiaoyueya/chain-sdk/common"
)

func AddOrUpdateChain(chainInfo *common.ChainInfo) error {
	if chainInfo == nil {
		return errors.New("nil chain info")
	}
	buf, err := json.Marshal(chainInfo)
	if err != nil {
		return err
	}

	key := common.BuildChainKey(chainInfo.Chain)

	return redisService.client.Set(key, buf, 0).Err()
}

func GetChain(chain string) (*common.ChainInfo, error) {
	key := common.BuildChainKey(chain)

	buf, err := redisService.client.Get(key).Bytes()

	if err != nil {
		return nil, err
	}
	chainInfo := new(common.ChainInfo)
	err = json.Unmarshal(buf, chainInfo)
	if err != nil {
		return nil, err
	}

	return chainInfo, nil
}
