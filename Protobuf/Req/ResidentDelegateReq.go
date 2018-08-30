package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//居民委托任务请求 1096
func ResidentDelegateReq(Uid int32) []byte {
	ResidentDelegateReq := &AutoMsg.ResidentDelegateReq{}
	Data,_ = proto.Marshal(ResidentDelegateReq)
	Param,_ := json.Marshal(ResidentDelegateReq)
	log :=Mgo.Log{
		Uid:   Uid,
		MsgId: 1096,
		Name:  "居民委托任务请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog();
	return SendRev.Send(1096,Data)
}
