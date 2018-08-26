package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type UnlockNpc struct {
	NpcId int32
}
//解锁npc
func UnlockNpcReq(Uid int32,UnlockNpc *UnlockNpc) []byte{
	UnlockNpcReq := &AutoMsg.UnlockNpcReq{
		NpcId: UnlockNpc.NpcId,
	}
	Data,_ := proto.Marshal(UnlockNpcReq)
	Param,_ := json.Marshal(LoadStaffReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 2025,
		Name: "解锁npc请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(2025,Data)

}
