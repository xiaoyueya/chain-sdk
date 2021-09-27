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

	key := common.BuildCoinKey(coinInfo.CustomName)

	if coinInfo.IsErc20 {
		redisService.client.SAdd(common.BucketCoinSetKey, coinInfo.CustomName)
		redisService.client.SAdd(coinInfo.GetContractSetKey(), coinInfo.CustomName)
	}

	return redisService.client.Set(key, buf, 0).Err()
}

func GetAllXrc20s() ([]string, error) {
	return redisService.client.SMembers(common.BucketCoinSetKey).Result()
}

// GetAllChainContracts 获取某链下所有的xrc20的合约地址
func GetAllChainCoinKeys(chain string) ([]string, error) {
	return redisService.client.SMembers(common.BuildChainContractsKey(chain)).Result()
}

func GetCoinOrToken(customName string) (*common.CoinOrToken, error) {
	key := common.BuildCoinKey(customName)

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
