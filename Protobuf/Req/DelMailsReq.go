package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type DelMails struct {
	Id []string
}
//批量删除邮件 1080
func DelMailsReq(Uid int32,DelMails *DelMails) []byte {
	DelMailsReq := &AutoMsg.DelMailsReq{
		Id: DelMails.Id,
	}
	Data,_ := proto.Marshal(DelMailsReq)
	Param,_ := json.Marshal(DelMailsReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1080,
		Name:  "批量删除邮件",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1080,Data)
}
