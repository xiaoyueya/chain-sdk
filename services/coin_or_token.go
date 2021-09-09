package services

import (
	"encoding/json"
	"errors"
	"github.com/xiaoyueya/chain-sdk/common"
)

func AddOrUpdateCoin(coinInfo *common.CoinOrToken) error {
	if coinInfo == nil {
		return errors.New("nil coin info")
	}
	buf, err := json.Marshal(coinInfo)
	if err != nil {
		return err
	}

	key := common.BuildCoinKey(coinInfo.Chain, coinInfo.ChainId, coinInfo.Symbol)

	return redisService.client.Set(key, buf, 0).Err()
}

func GetCoinOrToken(chain, chainId, symbol string) (*common.CoinOrToken, error) {
	key := common.BuildCoinKey(chain, chainId, symbol)

	buf, err := redisService.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	coinInfo := new(common.CoinOrToken)
	err = json.Unmarshal(buf, coinInfo)
	if err != nil {
		return nil, err
	}

	return coinInfo, nil
}
