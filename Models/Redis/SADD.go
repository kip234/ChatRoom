package Redis

import "fmt"

func (r RedisPool)SADD(args ...interface{}) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("SADD",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	return nil
}
