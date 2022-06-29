package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RedisConf  redis.RedisConf
	CacheRedis cache.CacheConf
	Mysql      struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	ElasticSearch struct {
		Addresses []string
		UserName  string
		PassWord  string
	}
}
