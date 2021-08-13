package main

import (
	"ChatRoom/Models/Data"
	"ChatRoom/client"
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	udata:=Data.Udata{}

	dict:=make(map[string]string)
	dict["-h"]="host"
	dict["-n"]="name"
	dict["-p"]="pwd"
	dict["-u"]="uid"
	dict["-r"]="-r"

	info:=client.Args(dict)
	//因为UID为int所以特地断言保证服务端可以正常赋值
	if _,ok:=info["uid"];ok{
		uid, err := strconv.Atoi(info["uid"].(string))
		if err != nil {
			panic(err)
		}
		info["uid"] = uid
	}

	var ok bool
	if info["-r"],ok=info["-r"].(string);!ok {
		fmt.Printf("The -r parameter is missing\n")
		return
	}
	//发送请求
	resp:=client.Send(info["host"].(string),info["-r"].(string),info)

	//解析返回的消息
	var m Data.Message
	err:=json.Unmarshal(resp,&m)
	if err!=nil {
		panic(err)
	}
	if m.Typ==Data.ErrTyp {
		fmt.Println(m.Content)
		return
	}

	err=json.Unmarshal([]byte(m.Content),&udata)
	if err!=nil {
		panic(err)
	}

	//如果是注册就直接结束
	if info["-r"]=="register" {
		fmt.Println(m.Content)
		return
	}

	//建立websocket链接
	conn:=client.Upgrad(info["host"].(string),udata.Token)

	//接收消息
	go client.Rec(conn)

	client.GetIn(&udata,conn)
}