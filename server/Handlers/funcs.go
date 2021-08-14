package Handlers

import (
	"ChatRoom/Models"
	"ChatRoom/Models/Data"
	"ChatRoom/Models/User"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
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
		content string,						//Message内容
		home,in *chan Data.Message,			//大厅信道,当前使用的信道
		rooms map[string]*Models.ConnPool,	//各个房间
		uid int,							//当前用户的UID
		conn *websocket.Conn,				//当前用户的链接
		blklsts map[int]*Models.BlockList,	//屏蔽名单
		sql *gorm.DB,						//SQL链接
		) error {

	cmd:=Data.Cmd{}
	err:=json.Unmarshal([]byte(content),&cmd)
	if err!=nil {
		return err
	}
	log.Printf("cmd: %v\n",cmd)
	switch cmd.Code{
	case Data.InRoom:
		for _,i:=range cmd.Content{
			if _,ok:=rooms[i];ok{
				*in=rooms[i].Chan()
				rooms[i].Add(uid,conn)
			}else {
				return fmt.Errorf("invalid room name %s",i)
			}
		}

	case Data.NwRoom:
		if len(cmd.Content)!=1{
			return fmt.Errorf("only one new room can be created at a time")
		}
		if cmd.Content[0]=="lobby"{
			return fmt.Errorf("illegal room name")
		}
		if _,ok:=rooms[cmd.Content[0]];ok {
			return fmt.Errorf("the room name already exists")
		}
		rooms[cmd.Content[0]]=Models.NewLobby(cmd.Content[0])
		*in=rooms[cmd.Content[0]].Chan()
		rooms[cmd.Content[0]].Add(uid,conn)
		go rooms[cmd.Content[0]].Run(true,blklsts)

	case Data.OutRoom:
		for _,i:=range cmd.Content{
			rooms[i].Del(uid)
			//空房间
			if rooms[i].ConnNum()==0{
				delete(rooms,i)
			}
		}
		*in=*home

	case Data.LstRoom:
		var keys []string
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

	case Data.Block:
		for _,i:=range cmd.Content{
			u:=User.User{}//因为UID为主键,这里放在循环内部是为了清理每次残留的信息避免主键不变
			sql=sql.Where(&User.User{Name: i}).Find(&u)//获取UID
			err=sql.Error
			blklsts[uid].Add(u.Uid,i)
		}

	case Data.UnBlk:
		for _,i:=range cmd.Content{
			u:=User.User{}
			sql=sql.Where(&User.User{Name: i}).Find(&u)//获取UID
			err=sql.Error
			blklsts[uid].UnBlock(u.Uid)
		}

	default:
		err = fmt.Errorf("unknown code")
	}
	return err
}