package User

import (
	"ChatRoom/Models/Redis"
	"gorm.io/gorm"
	"strconv"
)

func UserData(uid int) string {
	return strconv.Itoa(uid)+"udt"
}

func Out(db *gorm.DB,redis Redis.RedisPool) (err error) {
	var p []User
	db=db.Find(&p)
	if db.Error!=nil {
		return db.Error
	}
	for _,i:=range p{
		err=redis.HMSET(UserData(i.Uid),"name",i.Name,"pwd",i.Pwd) //
		if err!=nil {
			return err
		}
	}
	return nil
}
