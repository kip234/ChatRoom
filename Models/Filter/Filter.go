package Filter

import (
	"ChatRoom/Models/Filter/prefix_tree"
	"io/ioutil"
	"os"
	"strings"
)

type Filter struct{
	tree    *prefix_tree.Prefix_tree
	replace byte
}

func (f *Filter)Add(content []byte){
	f.tree.Add(content)
}

func (f *Filter)Process(content []byte) (ok bool,re []byte) {
	location:=f.tree.Find(content)
	index1:=0//content下标
	index2:=0//location下标
	length:=len(content)
	nums:=len(location)
	if nums == 0{//没有出现
		return false,content
	}

	for index1<length{
		if index2<nums && index1==location[index2][0]{//开始出现违规文字
			for index1<location[index2][1]{
				re=append(re,f.replace)
				if content[index1]>127{
					index1+=3//中文
				}else{
					index1+=1
				}
			}
			index2+=1
		}else{
			re=append(re,content[index1])
			index1+=1
		}
	}
	ok=true
	return
}

func NewFilter(name string) (error,*Filter) {
	file,err:=os.Open(name)
	if err!=nil{
		return err,nil
	}
	b,err:=ioutil.ReadAll(file)
	if err!=nil{
		return err,nil
	}
	s:=string(b)
	words:=strings.Split(s,"\r\n")

	f:=Filter{
		tree: prefix_tree.NewPrefix_tree(),
	}
	f.replace=[]byte(words[0])[0]
	i:=1
	nums:=len(words)
	for i<nums{
		f.Add([]byte(words[i]))
		i+=1
	}
	return nil,&f
}

