package Handlers

import (
	"ChatRoom/Models"
	"ChatRoom/Models/Data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Lobby(pool *Models.ConnPool,rooms map[string]*Models.ConnPool) gin.HandlerFunc {
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

				fmt.Println(m)

				if m.Typ==Data.UMgTyp{
					in<-m
				}else if m.Typ==Data.CmdTyp{
					err=CmdProc(string(m.Content),&home,&in,rooms,uid,conn)
					if err!=nil {
						conn.WriteJSON(Data.Message{
							Typ: Data.ErrTyp,
							Content:[]byte(err.Error()),
						})
					}
				}
			}
		}()
	}
}
