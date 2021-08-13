package main

import (
	"ChatRoom/Models"
	"ChatRoom/Models/JWT"
	"ChatRoom/Models/Redis"
	"ChatRoom/Models/User"
	"ChatRoom/server/Database"
	"ChatRoom/server/Routers"
	"ChatRoom/server/config"
)

const ConfPath 	= "conf.json"	//外部配置文件
const Secret 	= "I'mTooDishes"//JWT秘钥
var conf config.Conf

var (
	DefaultJwt = JWT.Jwt{
		Header: JWT.Header{
			Alg: "HS256",
			Typ: "JWT",
		},
		Payload: JWT.Payload{
			Iss: "kip",
			Sub: "ChatRoom",
		},
		Secret: Secret,
	}
)



func main()  {
	conf=config.Init(ConfPath)//获取服务器配置
	var redis = Redis.RedisPool{
		Write:		conf.Wredis,
		Read:		conf.Rredis,
		IdLeTimeout:5,
		MaxIdle:	20,
		MaxActive:	8,
	}

	redis.Init()
	db:=Database.InitGorm(&conf.Sql)
	lobby:=Models.NewLobby("lobby")
	go lobby.Run(false)//要一直等待连接
	//房间名
	rooms:=make(map[string]*Models.ConnPool)
	//redis.
	//缓存数据
	User.Out(db,redis)
	server:=Routers.BuildRouter(db,&redis,&DefaultJwt,lobby,rooms)
	server.Run(conf.Addr)
}
