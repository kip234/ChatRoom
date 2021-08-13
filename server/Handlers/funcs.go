package Handlers

import (
	"ChatRoom/Models"
	"ChatRoom/Models/Data"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//用于从上下文中取出UID
func getUid(c *gin.Context) (uid int,err error) {
	v,ok:=c.Get("uid")
	if !ok {
		uid=-1
		err = fmt.Errorf("Missing UID field")
		return
	}
	uid,ok = v.(int)
	if !ok {
		uid=-1
		err = fmt.Errorf("Assertion failure")
		return
	}
	return uid,nil
}


func CmdProc(
		content string,
		home,in *chan Data.Message,
		rooms map[string]*Models.ConnPool,
		uid int,
		conn *websocket.Conn,
		) error {

	cmd:=Data.Cmd{}
	err:=json.Unmarshal([]byte(content),&cmd)
	if err!=nil {
		return err
	}

	switch cmd.Code{
	case Data.InRoom:
		if _,ok:=rooms[cmd.Content];ok{
			*in=rooms[cmd.Content].Chan()
			rooms[cmd.Content].Add(uid,conn)
		}else {
			return fmt.Errorf("invalid room name")
		}
	case Data.NwRoom:
		if cmd.Content=="lobby"{
			return fmt.Errorf("illegal room name")
		}
		if _,ok:=rooms[cmd.Content];ok {
			return fmt.Errorf("the room name already exists")
		}
		rooms[cmd.Content]=Models.NewLobby(cmd.Content)
		*in=rooms[cmd.Content].Chan()
		rooms[cmd.Content].Add(uid,conn)
		go rooms[cmd.Content].Run(true)
	case Data.OutRoom:
		*in=*home
		rooms[cmd.Content].Del(uid)

		//空房间
		if rooms[cmd.Content].ConnNum()==0{
			delete(rooms,cmd.Content)
		}

	case Data.LstRoom:
		keys:=make([]string,len(rooms))
		for key,_:=range rooms{
			keys=append(keys,key)
		}
		c,err:=json.Marshal(keys)
		fmt.Sprintln(keys)
		if err!=nil {
			return err
		}

		m:=Data.Message{
			Typ: Data.SrMTyp,
			PoolName: "lobby",
			Content: c,
		}
		err=conn.WriteJSON(m)
	default:
		err = fmt.Errorf("unknown code")
	}
	return err
}