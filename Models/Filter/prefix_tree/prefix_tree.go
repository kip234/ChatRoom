package prefix_tree

//go的byte似乎是十进制ASCII码，传闻ASCII码阈值0x7f=126或者0xFF
type Prefix_tree struct {
	roots map[byte]*node
}

func NewPrefix_tree()*Prefix_tree{
	return &Prefix_tree{
		roots:make(map[byte]*node),
	}
}

//添一条记录
func (p *Prefix_tree)Add(str []byte){
	index:=0                //当前str处理的下标
	lenth:=len(str)         //str长度
	var now,proc *node		//当前指向的节点,要处理的节点
	flag:=false             //匹配状态

	//找根节点
	proc= newNode(str[index])
	if _,flag=p.roots[str[index]];flag{//有匹配的
		now=p.roots[str[index]]
	}

	if !flag{//没找到?
		p.roots[str[index]]=proc
		now=proc
	}
	index+=1

	for index<lenth{
		proc= newNode(str[index])
		flag=false
		if _,flag=now.Sons[str[index]];flag{
			now=now.Sons[str[index]]
		}

		//没有？
		if !flag {
			now.Sons[str[index]]=proc
			now=proc
		}
		index+=1
	}
}

//查找出现的记录
func (p *Prefix_tree)Find(str []byte) (re [][2]int) {
	index:=0                //当前str处理的下标
	lenth:=len(str)         //str长度
	var now *node 			//当前指向的节点,要处理的节点
	flag:=false             //匹配状态
	be:=index               //出现的位置

	for index<lenth{
		//找根节点
		flag=false
		if _,flag=p.roots[str[index]];flag{
			now=p.roots[str[index]]
			index+=1
		}
		if !flag{//没找到?
			index+=1
			be=index
			continue
		}
		//匹配后续字符
		for index<lenth{
			flag=false
			if _,flag=now.Sons[str[index]];flag{
				now=now.Sons[str[index]]
			}
			if !flag {//没有？
				be=index
				break
			}else if flag && len(now.Sons)==0{//到末尾了，成功匹配到一个
				index+=1
				re=append(re,[2]int{be,index})
				be=index
				break
			}
			index+=1
		}
	}
	return
}