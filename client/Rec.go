package client

import (
	"ChatRoom/Models/Data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func Rec(conn *websocket.Conn){
	defer conn.Close()
	for true {

		m:=Data.Message{}
		err:=conn.ReadJSON(&m)
		if err!=nil {
			fmt.Printf("Connection disconnect !\n")
			break
		}

		if m.Typ==Data.UMgTyp{
			um:=Data.UMg{}
			json.Unmarshal(m.Content,&um)
			fmt.Printf("[%s]%s:%s\n",m.PoolName,um.Name,um.Content)
		}else if m.Typ==Data.SrMTyp{
			fmt.Printf("[Server]:%s\n",m.Content)
		}else if m.Typ==Data.ErrTyp{
			fmt.Printf("[Error]:%s\n",m.Content)
		}
	}
}
