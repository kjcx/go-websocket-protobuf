package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type NpcRelationAdvance struct {
	NpcId int32
}
//提升居民品质 1037
func NpcRelationAdvanceReq(NpcRelationAdvance *NpcRelationAdvance) []byte{
	NpcRelationAdvanceReq := &AutoMsg.NpcRelationAdvanceReq{
		NpcId: NpcRelationAdvance.NpcId,
	}
	Data,_ :=proto.Marshal(NpcRelationAdvanceReq)
	return SendRev.Send(1037,Data)
}