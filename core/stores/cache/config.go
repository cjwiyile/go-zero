package cache

import "github.com/yilefreedom/go-zero/core/stores/redis"

type (
	ClusterConf []NodeConf

	NodeConf struct {
		redis.RedisConf
		Weight int `json:",default=100"`
	}
)
