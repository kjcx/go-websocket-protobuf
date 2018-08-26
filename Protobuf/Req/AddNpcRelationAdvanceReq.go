package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type AddNpcFavorability struct {
	NpcId int32
	ItemId int32
	ItemCount int32
}
//增加好感度
func AddNpcRelationAdvanceReq(AddNpcFavorability *AddNpcFavorability) []byte {
	favorability := &AutoMsg.AddNpcFavorability{
		NpcId:     AddNpcFavorability.NpcId,
		ItemId:    AddNpcFavorability.ItemId,
		ItemCount: AddNpcFavorability.ItemCount,
	}
	AddNpcRelationAdvanceReq := &AutoMsg.AddNpcRelationAdvanceReq{
		AddNpcFavor: favorability,
	}
	Data,_ := proto.Marshal(AddNpcRelationAdvanceReq)
	return SendRev.Send(1038,Data)
}