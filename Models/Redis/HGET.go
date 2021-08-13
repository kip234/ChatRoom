package Redis

import "fmt"

func (r RedisPool)HGET(key,field  string) (string,error) {
	rdb :=r.rpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("HGET",key,field)
	if err != nil{
		return "",err
	}
	if v==nil {
		return "",nil
	}
	re,ok:=v.([]uint8)
	if !ok {
		return "",fmt.Errorf("func (r RedisPool)HGET(key,field  string) (string,error) : Assertion failure")
	}
	return string(re),err
}
