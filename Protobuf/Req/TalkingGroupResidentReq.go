package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type TalkingGroupResident struct {
	Ids                  []int32
	Replace              bool
}
//谈判团居民任命请求 2031
func TalkingGroupResidentReq(Uid int32,TalkingGroupResident *TalkingGroupResident) []byte {
	TalkingGroupResidentReq := &AutoMsg.TalkingGroupResidentReq{
		Ids:     TalkingGroupResident.Ids,
		Replace: TalkingGroupResident.Replace,
	}
	Data,_ := proto.Marshal(TalkingGroupResidentReq)
	Param,_ := json.Marshal(TalkingGroupResidentReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2031,
		Name:  "谈判团居民任命请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2031,Data)
}
