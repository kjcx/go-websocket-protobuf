package Model

import (
	"fmt"
	"github.com/golang/protobuf/proto/test_proto"
	"WebSocket/Mgo"
	"WebSocket/Protobuf/Req"
	"time"
)

func main() {
	
}

type LoadSignList struct {
	Day int32
	IsSign bool
}
type LoadSignInfo struct {
	LoadSignList [35]*LoadSignList
}
type SignResult struct {
	Month int32
	LoaSignInfo *LoadSignInfo
}
//签到解析
func Sign(msg *AutoMsg.SignResult) (*SignResult,int32) {
	session := Mgo.Mongo()
	c := session.DB("test").C("Sign")
	SignData := &SignResult{}
	var IsSignDay int32
	day := int32(time.Now().Day())//今日
	for k,v := range msg.GetLoaSignInfo() {
		//fmt.Println(k)
		//fmt.Println(v.GetLoadSignList())
		list := &LoadSignInfo{}
		for kk,vv :=range v.GetLoadSignList() {
			if vv.GetDay() == day  {
				if !vv.GetIsSign(){
					IsSignDay = day
				}else {
					IsSignDay = 0
				}
			} else {
				IsSignDay = 0
			}
			info := &LoadSignList{vv.GetDay(), vv.GetIsSign()}
			//fmt.Println(info)
			list.LoadSignList[kk] = info
		}
		signInfo := &LoadSignInfo{list.LoadSignList}
		SignData = &SignResult{Month:k,LoaSignInfo:signInfo}
		//fmt.Println(SignData)
		err := c.Insert(SignData)
		if err!=nil{
			fmt.Println("Sign 插入失败")
		}else{
			fmt.Println("插入成功")
		}
	}
	return SignData,IsSignDay
}
//执行签到
func DaySign(Uid int32,Day int32) []byte{
	fmt.Println("执行今日签到:",Day)
	return Req.DaySignReq(Uid,Day)
}