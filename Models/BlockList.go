package Models

type BlockList struct{
	uid int//所属用户
	list map[int]string//被屏蔽的用户
}

//屏蔽某人
func (b *BlockList)Add(uid int,name string){
	b.list[uid]=name
}

//判断是否被屏蔽
func (b *BlockList)Blocked(uid int) (ok bool) {
	_,ok=b.list[uid]
	return
}

//解除屏蔽
func (b *BlockList)UnBlock(uid int){
	delete(b.list,uid)
}

func NewBlockList(uid int) *BlockList{
	return &BlockList{uid:uid,list:make(map[int]string)}
}