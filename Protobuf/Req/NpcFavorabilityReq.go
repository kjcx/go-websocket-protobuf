package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//加载npc好感度 1036
func NpcFavorabilityReq(Uid int32) []byte{
	NpcFavorabilityReq := &AutoMsg.NpcFavorabilityReq{}
	Data,_ := proto.Marshal(NpcFavorabilityReq)
	Param,_ := json.Marshal(NpcFavorabilityReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1036,
		Name: "加载npc好感度请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1036,Data)
}