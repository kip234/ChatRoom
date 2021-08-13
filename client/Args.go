package client

import (
	"os"
	"strings"
)

//获取命令行参数
func Args(dict map[string]string) map[string]interface{}{
	info:= make(map[string]interface{})
	for _,arg:= range os.Args{
		tmp:=strings.Split(arg,"=")
		if len(tmp)==2 {
			info[dict[tmp[0]]] = tmp[1]
		}
	}
	return info
}
