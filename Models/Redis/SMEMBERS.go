package Redis

import "fmt"

func (r RedisPool)SMEMBERS(key string) (re []string,err error) {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("SMEMBERS",key)
	fmt.Println(v)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)SMEMBERS(key string) ([]string,error): %s",err.Error())
	}
	tmp:=v.([]interface{})
	for _,i:=range tmp{
		re=append(re, string(i.([]uint8)))//这里很有可能会出问题
	}
	return re,nil
}
