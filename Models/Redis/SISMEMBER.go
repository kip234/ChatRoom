package Redis

import "fmt"

func (r RedisPool)SISMEMBER(args ...interface{}) (int64,error) {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("SISMEMBER",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	re,ok:=v.(int64)
	if !ok {
		return 0,fmt.Errorf("func (r RedisPool)SISMEMBER(args ...interface{}) (int,error) : Assertion error")
	}
	return re,nil
}
