package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//偷菜请求 1061
func StealSemenReq(Uid int32,StealSemen *AutoMsg.StealSemenReq) []byte {
	StealSemenReq := &AutoMsg.StealSemenReq{
		Id:     StealSemen.Id,
		UserId: StealSemen.UserId,
	}
	Data,_ := proto.Marshal(StealSemenReq)
	Param,_ := json.Marshal(StealSemenReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1061,
		Name:  "偷取",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1061,Data)
}
