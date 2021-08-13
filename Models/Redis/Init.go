package Redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func (r *RedisPool)Init()  {
	r.rpool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn,err := redis.Dial("tcp",r.Read)
			return conn,err
		},
		MaxIdle: r.MaxIdle,
		MaxActive: r.MaxActive,
		IdleTimeout: time.Second*time.Duration(r.IdLeTimeout),
	}

	r.wpool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn,err := redis.Dial("tcp",r.Write)
			return conn,err
		},
		MaxIdle: r.MaxIdle,
		MaxActive: r.MaxActive,
		IdleTimeout: time.Second*time.Duration(r.IdLeTimeout),
	}
}
