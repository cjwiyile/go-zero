package config

import (
	"github.com/yileCJW/go-zero/rest"
	"github.com/yileCJW/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf
	Check zrpc.RpcClientConf
}
