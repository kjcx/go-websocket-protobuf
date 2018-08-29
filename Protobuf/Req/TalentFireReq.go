package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
type TalentFire struct {
	RoleId string
}
//解雇经理请求
func TalentFireReq(Uid int32,TalentFire *TalentFire) []byte{
	TalentFireReq := &AutoMsg.TalentFireReq{
		RoleId: TalentFire.RoleId,
	}
	Data,_ := proto.Marshal(TalentFireReq)
	Param,_ := json.Marshal(TalentFireReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1132,
		Name:  "解雇经理请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1132,Data)
}
