package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)
//加载员工请求 1083
func LoadStaffReq(Uid int32) []byte{
	LoadStaffReq := &AutoMsg.LoadStaffReq{}
	Data,_ := proto.Marshal(LoadStaffReq)
	Param,_ := json.Marshal(LoadStaffReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1083,
		Name: "加载员工请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1083,Data)
}
