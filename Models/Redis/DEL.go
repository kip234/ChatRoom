package Redis

import "fmt"

func (r RedisPool)DEL(key string) error {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("DEL",key)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	return err
}
