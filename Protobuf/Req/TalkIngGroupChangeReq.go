package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type TalkIngGroupChange struct {
	UpId   string
	DownId string
}
//谈判团换任请求 1114
func TalkIngGroupChangeReq(Uid int32,TalkIngGroupChange *TalkIngGroupChange) []byte {
	TalkIngGroupChangeReq := &AutoMsg.TalkIngGroupChangeReq{
		UpId:   TalkIngGroupChange.UpId,
		DownId: TalkIngGroupChange.DownId,
	}
	Data,_ := proto.Marshal(TalkIngGroupChangeReq)
	Param,_ := json.Marshal(TalkIngGroupChangeReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1114,
		Name:  "谈判团换任请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1114,Data)
}
