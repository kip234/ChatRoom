package client

import (
	"ChatRoom/Models/Data"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"os"
	"strings"
)

func GetIn(udata *Data.Udata,conn *websocket.Conn){
	defer conn.Close()

	command:=make(map[string]int)
	command["inroom"]=Data.InRoom
	command["lstroom"]=Data.LstRoom
	command["outroom"]=Data.OutRoom
	command["nwroom"]=Data.NwRoom
	command["block"]=Data.Block
	command["unblk"]=Data.UnBlk

	reader:=bufio.NewReader(os.Stdin)

	for true{

		//获取键盘输入

		b,_:=reader.ReadBytes('\n')
		for string(b)=="\r\n"||string(b)==""{
			fmt.Printf("input>")
			b,_=reader.ReadBytes('\n')
		}

		s:=string(b)
		s=strings.Trim(s,"\r\n")
		if s=="#q"||s=="#exit"{
			break
		}

		m:=Data.Message{}

		if string(s[0])=="/" {
			s=s[1:]
			var code int
			cmd:=strings.Split(s," ")
			var ok bool
			if code,ok=command[cmd[0]];!ok{
				fmt.Printf("unknown command %s\n",cmd)
				continue
			}
			cm := Data.Cmd{Code:code}
			if len(cmd)>1 {
				cm.Content=cmd[1:]
			}
			b,_:=json.Marshal(cm)
			m.Typ=Data.CmdTyp
			m.Content=b
		}else{
			um:=Data.UMg{
				Uinfo:Data.Uinfo{
					Uid: udata.Uid,
					Name: udata.Name,
				},
				//Content:[]byte(s),
			}
			if file,err:=os.Open(s);err==nil{
				um.Content,err=ioutil.ReadAll(file)
				if err!=nil{
					panic(err)
				}
				m.Typ=Data.FlTyp
			}else{
				um.Content=[]byte(s)
				m.Typ=Data.UMgTyp
			}
			b,err:=json.Marshal(um)
			if err!=nil{
				panic(err)
			}
			m.Content=b
		}
		//发送
		conn.WriteJSON(m)
	}
}
