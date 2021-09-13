package util

import (
	"github.com/xiaoyueya/chain-sdk/bwlog"
	"runtime/debug"
)

func DeferFunc(eFunc func() error) {
	err := eFunc()
	if err != nil {
		bwlog.Logger.Errorf("defer func error=%v,stack=%s\n", err, debug.Stack())
	}
}
