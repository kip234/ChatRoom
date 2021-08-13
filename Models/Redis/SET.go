package Redis

import "fmt"

func (r RedisPool)SET(args ...interface{}) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("SET",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)SET(args ...interface{}) error: %s",err.Error())
	}
	return err
}
