package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)

type TalentHire struct {
	ShopId string
	RoleId string
}
//雇佣经理 1131
func TalentHireReq(Uid int32,TalentHire *TalentHire) []byte{
	TalentHireReq := &AutoMsg.TalentHireReq{
		ShopId: TalentHire.ShopId,
		RoleId: TalentHire.RoleId,
	}
	Data,_ := proto.Marshal(TalentHireReq)
	Param,_ := json.Marshal(TalentHireReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1131,
		Name:  "雇佣经理请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1131,Data)
}
