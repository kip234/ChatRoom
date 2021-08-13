package Handlers

import (
	"ChatRoom/Models"
	"ChatRoom/Models/Data"
	"ChatRoom/Models/JWT"
	"ChatRoom/Models/Redis"
	"ChatRoom/Models/User"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(redis *Redis.RedisPool,template *JWT.Jwt,lobby *Models.ConnPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)//获取UID
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			return
		}

		//阻止重复登录
		if lobby.IsExit(uid) {
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte("the user is online"),
			})
			return
		}

		name,err:=redis.HGET(User.UserData(uid),"name")
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			return
		}

		template.Payload.Aud=uid//更新Payload值
		token := template.Encoding()//计算理应正确的JWT
		ud:=Data.Udata{//构建结构
			Uinfo:Data.Uinfo{
				Uid:uid,
				Name: name,
			},
			Token: token,
		}

		data,err:=json.Marshal(ud)
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			return
		}

		redis.SET(strconv.Itoa(uid),token)//放入Redis
		c.JSON(http.StatusOK,gin.H{//把令牌返回给客户
			"typ":Data.UdTyp,
			"content":data,
		})
	}
}