package Models

import (
	"ChatRoom/Models/Data"
	"github.com/gorilla/websocket"
)

type ConnPool struct{
	num int
	name string//标识
	in chan Data.Message//等待输入消息
	links map[int] *websocket.Conn
}

func (c *ConnPool)Add(uid int,conn *websocket.Conn){
	c.links[uid]=conn
	c.num+=1
}

func (c *ConnPool)Del(uid int){
	delete(c.links,uid)
	c.num-=1
}

func (c *ConnPool)Chan() chan Data.Message{
	return c.in
}

func (c *ConnPool)IsExit(uid int) (ok bool) {
	_,ok=c.links[uid]
	return
}

func (c *ConnPool)Run(AutoDown bool){
	for {
		if AutoDown&&len(c.links)==0{//没有连接的时候自动退出
			break
		}
		select {
		case data := <-c.in:
			data.PoolName=c.name//加上标识
			for uid,i:=range c.links{
				err:=i.WriteJSON(data)
				if err!=nil {//认为用户断开
					c.Del(uid)
				}
			}
		}
	}
}

func (c *ConnPool)ConnNum()int{
	return c.num
}

func NewLobby(Name string) *ConnPool {
	return &ConnPool{
		num: 0,
		name: Name,
		in:make(chan Data.Message),
		links:make(map[int] *websocket.Conn),
	}
}