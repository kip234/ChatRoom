package client

import (
	"ChatRoom/Models/Data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"strconv"
	"time"
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

		//fmt.Printf("type:%d\n",m.Typ)

		switch m.Typ{
		case Data.UMgTyp:
			um:=Data.UMg{}
			json.Unmarshal(m.Content,&um)
			fmt.Printf("[%s]%s:%s\n",m.PoolName,um.Name,um.Content)
		case Data.SrMTyp:
			fmt.Printf("[Server]:%s\n",m.Content)
		case Data.ErrTyp:
			fmt.Printf("[Error]:%s\n",m.Content)
		case Data.FlTyp:
			um:=Data.UMg{}
			json.Unmarshal(m.Content,&um)
			name:=strconv.FormatInt(time.Now().Unix(),16)//"随机"生成文件名
			file,err:=os.Create(name)
			fmt.Printf("%s\n",file.Name())
			if err!=nil{
				panic(err)
			}
			file.Write(um.Content)
			file.Close()
			fmt.Printf("[%s]%s: MAY HAVE SENT A PICTURE\n",m.PoolName,um.Name)
		default:
			fmt.Printf("a message of unknown type was received\n %v\n",m)
		}
	}
}
