package Data

const (
	ErrTyp	=1000	//错误信息
	SrMTyp	=1001	//来自服务端的消息
	UMgTyp	=1002	//某用户发送的消息
	UdTyp	=1003	//用户信息
	RLtTyp	=1005	//房间列表
	CmdTyp	=1006	//客户端的特殊指令
)

const (
	InRoom	=2000	//进入房间
	NwRoom	=2001	//创建房间
	OutRoom	=2002	//离开房间
	LstRoom	=2003	//房间列表
)