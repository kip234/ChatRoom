package Redis

import "fmt"

func (r RedisPool)HMSET(args ...interface{}) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("HMSET",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)HMSET(args ...interface{}) error: %s",err.Error())
	}
	return err
}