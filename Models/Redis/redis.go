package Redis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisPool struct {
	Read 		string
	Write 		string
	IdLeTimeout	int
	MaxIdle		int
	MaxActive	int
	rpool *redis.Pool//负责读取
	wpool *redis.Pool//负责写入
}