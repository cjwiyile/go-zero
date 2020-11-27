package config

import (
	"github.com/yilefreedom/go-zero/rest"
	"github.com/yilefreedom/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Transform zrpc.RpcClientConf
}
