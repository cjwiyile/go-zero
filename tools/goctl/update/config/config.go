package config

import "github.com/yilefreedom/go-zero/core/logx"

type Config struct {
	logx.LogConf
	ListenOn string
	FileDir  string
	FilePath string
}
