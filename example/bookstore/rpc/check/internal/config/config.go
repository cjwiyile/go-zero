package config

import (
	"github.com/yileCJW/go-zero/core/stores/cache"
	"github.com/yileCJW/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
}
