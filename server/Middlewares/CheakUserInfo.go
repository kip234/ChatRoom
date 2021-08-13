package Middlewares

import (
	"ChatRoom/Models/Data"
	"ChatRoom/Models/Redis"
	"ChatRoom/Models/User"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheakUserInfo(redi *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User.User
		err:=c.ShouldBind(&user)
		if err!=nil{//绑定失败
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			c.Abort()
			return
		}
		tmp:= User.User{}
		if err:=tmp.LoadRedis(redi,user.Uid);err!=nil{//查找失败
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			c.Abort()
			return
		}
		if tmp.Pwd!=user.Pwd{//密码对不上
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte("password wrong !"),
			})
			c.Abort()
			return
		}
		c.Set("uid",tmp.Uid)//验证通过,存入UID
	}
}