package Handlers

import (
	"ChatRoom/Models"
	"ChatRoom/Models/Data"
	"ChatRoom/Models/Filter"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Lobby(
		pool *Models.ConnPool,				//emmm...用户连接池?
		rooms map[string]*Models.ConnPool,	//房间
		blklsts map[int]*Models.BlockList,	//屏蔽名单
		sql *gorm.DB,						//SQL链接
		Filter *Filter.Filter,				//敏感词过滤器
		) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:= getUid(c)
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":err.Error(),
			})
			return
		}

		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"typ": Data.ErrTyp,
				"content": []byte(err.Error()),
			})
			return
		}
		conn.WriteJSON(Data.Message{
			Typ: Data.SrMTyp,
			Content: []byte("hello?"),
		})

		pool.Add(uid,conn)//添加链接
		blklsts[uid]=Models.NewBlockList(uid)//创建黑名单

		go func() {//接收消息
			in:=pool.Chan()//发送消息的链接
			home:=in//保存链接
			m:=Data.Message{}
			for true {

				err:=conn.ReadJSON(&m)
				if err!=nil {
					conn.Close()
					pool.Del(uid)
					break
				}

				//fmt.Printf("%v\n",m)

				if m.Typ==Data.CmdTyp{
					err=CmdProc(string(m.Content),&home,&in,rooms,uid,conn,blklsts,sql)
					if err!=nil {
						conn.WriteJSON(Data.Message{
							Typ: Data.ErrTyp,
							Content:[]byte(err.Error()),
						})
					}
				}else{
					if m.Typ==Data.UMgTyp{
						umg:=Data.UMg{}
						json.Unmarshal(m.Content,&umg)
						fmt.Printf("%s\n",string(umg.Content))
						_,umg.Content=Filter.Process(umg.Content)//过滤一下
						fmt.Printf("%s\n",string(umg.Content))
						m.Content,_=json.Marshal(umg)
					}
					in<-m
				}
			}
		}()
	}
}
