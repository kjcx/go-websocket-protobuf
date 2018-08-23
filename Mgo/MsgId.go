package Mgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"strings"
)

type MsgInfo struct {
	MsgId int32 //消息id
	Name string //结构体名字
	Notes string //注释
}
//定义一个数组接收数据
type AllInfo struct {
	List []*MsgInfo
}
//插上mgsid 表
func MsgProto(session *mgo.Session,info *MsgInfo) {
	c := session.DB("test").C("MsgId")
	c.Insert(info)
}

func GetMsg(session *mgo.Session,MsgId int32) string{
	c := session.DB("test").C("MsgId")
	//创建数组
	//通过单个数据结构创建数据接收数据
	AllInfo := make([]MsgInfo,0,10)
	//AllInfo := &AllInfo{}
	err:= c.Find(bson.M{"msgid":MsgId}).All(&AllInfo)
	fmt.Println(err)
	var ResultName = ""
	for _,v := range AllInfo {
		if strings.HasSuffix(v.Name,"Result"){
			ResultName = v.Name
		}
	}
	return ResultName
}