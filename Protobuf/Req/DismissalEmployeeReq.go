package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
)

type DismissalEmployee struct {
	ListId []string
}
//解雇员工请求 1087
func DismissalEmployeeReq(Uid int32,DismissalEmployee *DismissalEmployee)[]byte  {
	DismissalEmployeeReq := &AutoMsg.DismissalEmployeeReq{
		ListId: DismissalEmployee.ListId,
	}
	Data,_ := proto.Marshal(DismissalEmployeeReq)
	Param,_ := json.Marshal(DismissalEmployeeReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1087,
		Name:  "解雇员工请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1087,Data)

}
