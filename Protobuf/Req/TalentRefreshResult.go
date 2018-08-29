package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)
//刷新人才市场1135
func TalentRefreshReq(Uid int32) []byte{
	TalentRefreshReq := &AutoMsg.TalentRefreshReq{}
	Data,_ := proto.Marshal(TalentRefreshReq)
	Param,_ := json.Marshal(TalentRefreshReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1135,
		Name:  "刷新人才市场",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1135,Data)
}
