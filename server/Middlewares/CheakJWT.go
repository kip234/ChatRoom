package Middlewares

import (
	"ChatRoom/Models/Data"
	"ChatRoom/Models/JWT"
	"ChatRoom/Models/Redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//读取
func CheakJWT(pool *Redis.RedisPool,template *JWT.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")//找token
		if token == ""{//没有token
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte("token == nil"),
			})
			c.Abort()
			return
		}

		err:=template.Decoding(token)//解码-读取Aud
		if err!=nil{//有问题
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			c.Abort()
			return
		}

		s,err:=pool.GET(strconv.Itoa(template.Payload.Aud))//查找Redis记录
		if err!=nil{//有问题
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			c.Abort()
			return
		}

		if s!=token{//和记录的令牌对不上
			c.JSON(http.StatusOK,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte("token error"),
			})
			c.Abort()
			return
		}

		c.Set("uid",template.Payload.Aud)//存入UID
	}
}