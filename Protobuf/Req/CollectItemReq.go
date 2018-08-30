package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type CollectItem struct {
	RoleId int32
	ItemId []int32
}
//商战胜利收取道具请求 2046
func CollectItemReq(Uid int32,CollectItem *CollectItem) []byte {
	CollectItemReq := &AutoMsg.CollectItemReq{
		RoleId: CollectItem.RoleId,
		ItemId: CollectItem.ItemId,
	}
	Data,_ := proto.Marshal(CollectItemReq)
	Param,_ := json.Marshal(CollectItemReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2046,
		Name:  "商战胜利收取道具请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2046,Data)

}
