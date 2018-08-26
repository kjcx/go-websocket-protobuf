package Req

import (
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/golang/protobuf/proto/test_proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)


//结构体
type UseItem struct {
	ItemId int32
	Count int32
}
//使用道具 1015
func UseItemReq(Uid int32,UseItem *UseItem) []byte{
	//data := &UseItem{10010,1}
	UseItemReq := &AutoMsg.UseItemReq{
		ItemId:UseItem.ItemId,
		Count:UseItem.Count,
	}
	Data,err := proto.Marshal(UseItemReq)
	if err!=nil{
		log.Fatal("UseItemReq解析失败")
	}
	Param,_ := json.Marshal(UseItemReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1015,
		Name: "使用道具请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1015,Data)
}
