//暂时没弄
package Handlers

import (
	"ChatRoom/Models/Data"
	"ChatRoom/Models/Redis"
	"ChatRoom/Models/User"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//注册成狗后会返回UID
func Register(Sql *gorm.DB,redi *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User.User
		err:=c.ShouldBind(&user)//获取用户相关信息
		if err!=nil {//绑定出错
			c.JSON(http.StatusBadRequest,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			return
		}
		err=user.SaveMySQL(Sql)
		if err!=nil{//存入MySQL出错
			c.JSON(http.StatusBadRequest,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			return
		}

		err=user.SaveRedis(redi)
		if err!=nil{//存入Redis出错
			c.JSON(http.StatusBadRequest,gin.H{
				"typ":Data.ErrTyp,
				"content":[]byte(err.Error()),
			})
			return
		}

		//反馈
		ud:=Data.Udata{
			Uinfo: Data.Uinfo{
				Uid:user.Uid,
				Name:user.Name,
			},
		}
		p,_:=json.Marshal(ud)
		c.JSON(http.StatusOK,gin.H{
			"typ":Data.UdTyp,
			"content":p,
		})
	}
}

