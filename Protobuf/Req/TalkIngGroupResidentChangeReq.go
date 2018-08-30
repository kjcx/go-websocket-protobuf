package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type TalkIngGroupResidentChange struct {
	UpId                 int32
	DownId               int32
}
//更改居民请求 2033
func TalkIngGroupResidentChangeReq(Uid int32,TalkIngGroupResidentChange *TalkIngGroupResidentChange) []byte{
	TalkIngGroupResidentChangeReq := &AutoMsg.TalkIngGroupResidentChangeReq{
		UpId:   TalkIngGroupResidentChange.UpId,
		DownId: TalkIngGroupResidentChange.DownId,
	}
	Data,_ := proto.Marshal(TalkIngGroupResidentChangeReq)
	Param,_ := json.Marshal(TalkIngGroupResidentChangeReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2033,
		Name:  "更改居民请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2033,Data)

}
