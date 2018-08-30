package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type ReadMail struct {
	Id string
}
//阅读邮件请求 1081
func ReadMailReq(Uid int32,ReadMail *ReadMail) []byte  {
	ReadMailReq := &AutoMsg.ReadMailReq{
		Id: ReadMail.Id,
	}
	Data,_ := proto.Marshal(ReadMailReq)
	Param,_ := json.Marshal(ReadMailReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1081,
		Name:  "阅读邮件请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1081,Data)
}
