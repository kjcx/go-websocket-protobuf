package Mgo

import (
	"gopkg.in/mgo.v2"
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
func InsertLog(session *mgo.Session, Data *Log){
	mongo := session.DB("test").C("GoLog");
	mongo.Insert(Data)
}