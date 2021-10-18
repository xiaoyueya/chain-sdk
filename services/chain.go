package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xiaoyueya/chain-sdk/bwlog"
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

func AvailableClient(info *common.ChainInfo) (*ethclient.Client, error) {
	client, err := ethclient.Dial(info.RPCUrls[info.RpcIndex])
	if err != nil {
		return nil, err
	}
	info.RpcIndex = info.RpcIndex + 1
	if info.RpcIndex >= len(info.RPCUrls) {
		info.RpcIndex = 0
	}

	height, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	err = AddOrUpdateChain(info)
	if err != nil {
		return nil, err
	}
	if height == 0 {
		return nil, errors.New("rpc url invalid")
	}
	bwlog.Logger.Info("got available client success,rpc index = %d,block height=%d\n", info.RpcIndex, height)
	return client, err
}
