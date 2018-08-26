package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type AddNpcFavorability struct {
	NpcId int32
	ItemId int32
	ItemCount int32
}
//增加好感度
func AddNpcRelationAdvanceReq(Uid int32,AddNpcFavorability *AddNpcFavorability) []byte {
	favorability := &AutoMsg.AddNpcFavorability{
		NpcId:     AddNpcFavorability.NpcId,
		ItemId:    AddNpcFavorability.ItemId,
		ItemCount: AddNpcFavorability.ItemCount,
	}
	AddNpcRelationAdvanceReq := &AutoMsg.AddNpcRelationAdvanceReq{
		AddNpcFavor: favorability,
	}
	Data,_ := proto.Marshal(AddNpcRelationAdvanceReq)
	Param,_ := json.Marshal(AddNpcRelationAdvanceReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1038,
		Name: "增加好感度请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1038,Data)
}