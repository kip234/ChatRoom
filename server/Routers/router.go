//建立路由结构

package Routers

import (
	"ChatRoom/Models"
	"ChatRoom/Models/JWT"
	"ChatRoom/Models/Redis"
	"ChatRoom/server/Handlers"
	"ChatRoom/server/Middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildRouter(db *gorm.DB,
		pool *Redis.RedisPool,
		template *JWT.Jwt,
		lobby *Models.ConnPool,
		rooms map[string]*Models.ConnPool,
	) *gin.Engine {
	server:=gin.Default()

	group:=server.Group("/", Middlewares.CheakJWT(pool,template))
	{
		group.GET("/lobby", Handlers.Lobby(lobby,rooms))
	}

	server.POST("/register", Handlers.Register(db,pool))

	server.POST("/login", Middlewares.CheakUserInfo(pool),Handlers.Login(pool,template,lobby))

	return server
}
