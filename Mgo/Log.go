package Mgo

import (
	"time"
)


type Log struct {
	Uid int32 //用户id
	MsgId int32 //消息id
	Name string //结构体名称
	Param string //参数
	Date time.Time //当前时间
	Msg string //错误信息
}
//插入日志
func (Data *Log)InsertLog(){
	mongo := Mongo()
	c := mongo.DB("test").C("GoLog");
	c.Insert(Data)
}

type LogResult struct {
	MsgId int32
	Data []byte
}

