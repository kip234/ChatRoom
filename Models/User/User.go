package User

import (
	"ChatRoom/Models/Redis"
	//"ChatRoom/Models/User"
	"gorm.io/gorm"
)

type User struct {
	Uid int `json:"uid" gorm:"primaryKey"`
	Name string	`json:"name" gorm:"string not null"`//用户名
	Pwd string `json:"pwd" gorm:"string not null"`//用户密码
}

//创建用户
func (u *User)SaveMySQL(db *gorm.DB) (err error) {
	err=db.Create(u).Error
	return
}
//储存信息至Redis
func (u *User)SaveRedis(db *Redis.RedisPool) (err error) {
	err=db.HMSET(UserData(u.Uid),"name",u.Name,"pwd",u.Pwd)
	//err=db.Create(u).Error
	return
}

//根据提供的UID读取用户信息
//主要用于密码比对
func (u *User)LoadMySQL(db *gorm.DB,uid int) (err error) {
	err=db.Where("Uid=?",uid).Find(u).Error
	return
}

func (u *User)LoadRedis(db *Redis.RedisPool,uid int) (err error) {
	//err=db.Where("Uid=?",uid).Find(u).Error
	u.Uid=uid
	u.Name,err=db.HGET(UserData(uid),"name")
	if err!=nil {
		return
	}
	u.Pwd,err=db.HGET(UserData(uid),"pwd")
	return
}

//判断密码是否正确，如果不正确返回false
func (u *User)PwdIsRight(db *Redis.RedisPool) bool {
	pwd,err:=db.HGET(UserData(u.Uid),"pwd")
	if err!=nil{
		return false
	}
	return pwd==u.Pwd
}

//判断是否存在，如果不存在返回false
func (u *User)IsExist(db *gorm.DB) bool {
	tmp:= User{}
	db.Where("Uid=?", u.Uid).Find(&tmp)
	if nil == db.Error {
			return true
		}
	return false
}
