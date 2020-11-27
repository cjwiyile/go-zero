package config

import (
	"github.com/yilefreedom/go-zero/core/stores/cache"
	"github.com/yilefreedom/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
}
