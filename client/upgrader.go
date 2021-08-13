package client

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func Upgrad(host,token string) *websocket.Conn {
	h:=make(http.Header)
	h.Add("token",token)
	ws,_,err:=websocket.DefaultDialer.Dial("ws://"+host+"/lobby",h) //协议改变！！！！
	if err!=nil {
		panic(err)
	}
	return ws
}