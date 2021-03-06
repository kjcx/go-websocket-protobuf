package Req

import (
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
	"github.com/golang/protobuf/proto/test_proto"
)
type CultivateEmployee struct {
	ListId []string
}
//培训员工请求 1086
func CultivateEmployeeReq(Uid int32,CultivateEmployee *CultivateEmployee) []byte{
	CultivateEmployeeReq := &AutoMsg.CultivateEmployeeReq{
		ListId: CultivateEmployee.ListId,
	}
	Data,_ := proto.Marshal(CultivateEmployeeReq)
	Param,_ := json.Marshal(CultivateEmployeeReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1086,
		Name:  "培训员工请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1086,Data)
}
