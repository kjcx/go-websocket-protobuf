package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type NpcRelationAdvance struct {
	NpcId int32
}
//提升居民品质 1037
func NpcRelationAdvanceReq(Uid int32,NpcRelationAdvance *NpcRelationAdvance) []byte{
	NpcRelationAdvanceReq := &AutoMsg.NpcRelationAdvanceReq{
		NpcId: NpcRelationAdvance.NpcId,
	}
	Data,_ :=proto.Marshal(NpcRelationAdvanceReq)
	Param,_ := json.Marshal(NpcRelationAdvanceReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1037,
		Name: "提升居民品质请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1037,Data)
}