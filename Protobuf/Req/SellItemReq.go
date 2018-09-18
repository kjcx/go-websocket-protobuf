package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type SellItem struct {
	ItemId int32
	Count int32
}
//出售道具 1014
func SellItemReq(Uid int32,SellItem *SellItem) []byte{
	SellItemReq := &AutoMsg.SellItemReq{
		ItemId: SellItem.ItemId,
		Count:  SellItem.Count,
	}
	Data,_ := proto.Marshal(SellItemReq)
	Param,_ := json.Marshal(SellItemReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1014,
		Name: "出售道具请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1014,Data)
}
