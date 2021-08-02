package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		SqlDns string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
