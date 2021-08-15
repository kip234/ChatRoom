# guide

## 一些说明

> words.sensitive保存敏感词组，一行一组，第一行第一个字节视为用于替换的字节

## packge

### client

> 含客户端需要调用的函数

| 原型                                                     | 描述                          |
| -------------------------------------------------------- | ----------------------------- |
| func Args(dict map[string]string) map[string]interface{} | 获取client启动时的命令行参数  |
| func GetIn(udata *Data.Udata,conn *websocket.Conn)       | 负责获取控制台输入并发送      |
| func Rec(conn *websocket.Conn)                           | 接收服务器的消息并显示        |
| func Send(to,method string,m interface{}) []byte         | 建立websocket之前进行信息发送 |
| func Upgrad(host,token string) *websocket.Conn           | 更新协议获取websocket连接     |

### Models

> 含使用的所有数据结构

#### Data

> 用于服务器与客户端信息交互的结构，交互始终以Message结构为载体，类型代号在default.go内定义

##### Message

```go
type Message struct {
	Typ			int		`json:"typ" binding:"required"`		//消息类型
	Owner		int		`json:"owner" binding:"required"`	//所有者
	PoolName	string	`json:"poolname"`				   //来自某个群？
	Content		[]byte 	`json:"content" binding:"required"`	//消息内容
}
```

##### Uinfo

```go
type Uinfo struct {
	Uid int			`json:"uid"`
	Name string		`json:"name"`
}
```

##### UMg

> 用户发送的普通消息，建立websocket后使用

```go
type UMg struct {
	Uinfo
	Content []byte	`json:"content"`
}
```

##### Udata

> 用户信息，用于服务器返回登录信息给客户端

```go
type Udata struct{
	Uinfo
	Token string `json:"token"`
}
```

##### Cmd

> 来自客户端的功能性消息

```go
type Cmd struct{
	Code 	int			`json:"code"`	//功能代号
	Content []string	`json:"content"`//参数、内容
}
```

##### default.go

```go
//Message Typ 值
const (
	ErrTyp	=1000	//错误信息
	SrMTyp	=1001	//来自服务端的消息
	UMgTyp	=1002	//某用户发送的消息
	UdTyp	=1003	//用户信息
	FlTyp	=1004	//以文件的形式储存该类型
	CmdTyp	=1005	//客户端的特殊指令
)

//Cmd Code 值
const (
	InRoom	=2000	//进入房间
	NwRoom	=2001	//创建房间
	OutRoom	=2002	//离开房间
	LstRoom	=2003	//房间列表
	Block  	=2004	//屏蔽某用户
	UnBlk	=2005	//解除屏蔽
)
```

#### Filter

> 敏感词条过滤器，同时支持中英文

##### Filter.go

```go
type Filter struct{
	tree    prefix_tree.Prefix_tree
	replace byte
}
```

##### prefix_tree

> 一颗前缀树

###### Prefix_tree

> 前缀树的定义

```go
type Prefix_tree struct {
	roots []*node
}
```

###### node

> 前缀树的节点

```go
type node struct {
	unit byte
	Sons []*node
}
```

#### JWT

> json web token 的一些东西

#### Redis

> 操作Redis的一些东西

#### User

> 服务器和SQL交互的结构

#### BlockList

> 屏蔽列表？

```go
type BlockList struct{
	uid int//所属用户
	list map[int]string//被屏蔽的用户
}
```

#### connpool

> 连接池？从in接收数据发送给links

```go
type ConnPool struct{
	num int
	name string//标识
	in chan Data.Message//等待输入消息
	links map[int] *websocket.Conn
}
```

### server

> 服务器用到的东西

#### config

> 处理部分配置信息

#### Database

> 处理SQL操作

#### Handlers

> 路由函数

##### login

> 登录

##### register

> 注册

##### lobby

> 成功登录后客户端会自动在该路由下与服务器建立websocket连接

#### Middlewares

> 一些中间件

##### CheakJWT

> 验证JWT

##### CheakUserInfo

> 登录时验证信息

#### Routers

> 路由设置

```go
group:=server.Group("/", Middlewares.CheakJWT(pool,template))
{
	group.GET("/lobby", Handlers.Lobby(lobby,rooms,blklsts,db,Filter))
}
server.POST("/register", Handlers.Register(db,pool))
server.POST("/login", Middlewares.CheakUserInfo(pool),Handlers.Login(pool,template,lobby))
```

