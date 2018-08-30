package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type DownlineLog struct {
	RoleId int32
}
//下线带来奖励日志请求2041
func DownlineLogReq(Uid int32,DownlineLog *DownlineLog) []byte {
	DownlineLogReq := &AutoMsg.DownlineLogReq{
		RoleId: DownlineLog.RoleId,
	}
	Data,_ := proto.Marshal(DownlineLogReq)
	Param,_ := json.Marshal(DownlineLogReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2041,
		Name:  "下线带来奖励日志请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2041,Data)
}
